package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

// CreateEntry create the permissions
func (p Permissions) CreateEntry() (Permissions, error) {
	s, err := session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("DB_REGION")),
		Endpoint: aws.String(os.Getenv("DB_ENDPOINT")),
	})
	if err != nil {
		return Permissions{}, fmt.Errorf("create entry new session err: %w", err)
	}
	perms, err := convertPermissionsToDynamo(p.Permissions, p.Identifier)
	if err != nil {
		return Permissions{}, fmt.Errorf("create convert permissions err: %w", err)
	}
	svc := dynamodb.New(s)
	item := map[string]*dynamodb.AttributeValue{
		"identifier": {
			S: aws.String(p.Identifier),
		},
		"permissions": &perms,
	}

	input := &dynamodb.PutItemInput{
		TableName:           aws.String(os.Getenv("DB_TABLE")),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(#IDENTIFIER)"),
		ExpressionAttributeNames: map[string]*string{
			"#IDENTIFIER": aws.String("identifier"),
		},
	}
	_, putErr := svc.PutItem(input)
	if putErr != nil {
		if awsErr, ok := putErr.(awserr.Error); ok {
			switch awsErr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				return Permissions{}, fmt.Errorf("permission identifier already exists: %w", awsErr.Code())
			case "ValidationException":
				fmt.Println(fmt.Sprintf("validation err reason: %v", input))
				return Permissions{}, fmt.Errorf("validation error: %w", awsErr)
			default:
				fmt.Println(fmt.Sprintf("unknown code err reason: %v", input))
				return Permissions{}, fmt.Errorf("unknown code err: %w", awsErr)
			}
		} else {
			return Permissions{}, fmt.Errorf("unknown err: %w", putErr)
		}
	}

	return convertDynamoToPermission(item)
}

// RetrieveEntry get the permissions
func (p Permissions) RetrieveEntry() (Permissions, error) {
	s, err := session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("DB_REGION")),
		Endpoint: aws.String(os.Getenv("DB_ENDPOINT")),
	})
	if err != nil {
		return Permissions{}, fmt.Errorf("retrieve entry session err: %w", err)
	}
	svc := dynamodb.New(s)
	input := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("DB_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"identifier": {
				S: aws.String(p.Identifier),
			},
		},
	}
	result, err := svc.GetItem(input)
	if err != nil {
		return Permissions{}, fmt.Errorf("retrieve get item err: %w", err)
	}

	if result.Item == nil {
		return Permissions{
			Identifier: p.Identifier,
			Status:     "no permissions",
		}, nil
	}
	return convertDynamoToPermission(result.Item)
}

// UpdateEntry alter the permissions
func (p Permissions) UpdateEntry(n Permissions) (Permissions, error) {
	s, err := session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("DB_REGION")),
		Endpoint: aws.String(os.Getenv("DB_ENDPOINT")),
	})
	if err != nil {
		return Permissions{}, fmt.Errorf("update entry session err: %w", err)
	}
	perms, err := convertPermissionsToDynamo(n.Permissions, n.Identifier)
	if err != nil {
		return Permissions{}, fmt.Errorf("update entry convert permissions err: %w", err)
	}
	svc := dynamodb.New(s)
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("DB_TABLE")),
		ExpressionAttributeNames: map[string]*string{
			"#PERMISSIONS": aws.String("permissions"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":permissions": &perms,
		},
		Key: map[string]*dynamodb.AttributeValue{
			"identifier": {
				S: aws.String(p.Identifier),
			},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String("SET #PERMISSIONS = :permissions"),
	}
	ret, err := svc.UpdateItem(input)
	if err != nil {
		return Permissions{}, fmt.Errorf("update item err: %w", err)
	}

	return convertDynamoToPermission(ret.Attributes)
}

// DeleteEntry remove the permissions
func (p Permissions) DeleteEntry() (Permissions, error) {
	s, err := session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("DB_REGION")),
		Endpoint: aws.String(os.Getenv("DB_ENDPOINT")),
	})
	if err != nil {
		return Permissions{}, fmt.Errorf("delete entry session err: %w", err)
	}
	svc := dynamodb.New(s)
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(os.Getenv("DB_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"identifier": {
				S: aws.String(p.Identifier),
			},
		},
	}
	_, err = svc.DeleteItem(input)
	if err != nil {
		return Permissions{}, fmt.Errorf("delete item err: %w", err)
	}

	return Permissions{
		Identifier: p.Identifier,
		Status:     "deleted",
	}, nil
}

// ScanEntries get all the permisisons
func ScanEntries() ([]Permissions, error) {
	s, err := session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("DB_REGION")),
		Endpoint: aws.String(os.Getenv("DB_ENDPOINT")),
	})
	if err != nil {
		return []Permissions{}, fmt.Errorf("scan entries session err: %w", err)
	}
	svc := dynamodb.New(s)
	input := &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("DB_TABLE")),
	}
	result, err := svc.Scan(input)
	if err != nil {
		return []Permissions{}, fmt.Errorf("scan entries err: %w", err)
	}
	itemLen := len(result.Items)
	if itemLen >= 1 {
		perms := []Permissions{}
		for i := 0; i < itemLen; i++ {
			item := result.Items[i]
			perm, err := convertDynamoToPermission(item)
			if err != nil {
				return perms, fmt.Errorf("couldnt convert perm: %w", err)
			}
			perms = append(perms, perm)
		}

		return perms, nil
	}

	return []Permissions{}, nil
}

func convertPermissionsToDynamo(perms []Permission, ident string) (dynamodb.AttributeValue, error) {
	ret := dynamodb.AttributeValue{}
	lMap := []*dynamodb.AttributeValue{}

	if len(perms) >= 1 {
		for _, perm := range perms {
			identifier := perm.Identifier
			if perm.Identifier == "" {
				identifier = ident
			}

			retMap := map[string]*dynamodb.AttributeValue{}
			retMap["name"] = &dynamodb.AttributeValue{
				S: aws.String(perm.Name),
			}
			retMap["action"] = &dynamodb.AttributeValue{
				S: aws.String(perm.Action),
			}
			retMap["identifier"] = &dynamodb.AttributeValue{
				S: aws.String(identifier),
			}
			mmap := &dynamodb.AttributeValue{
				M: retMap,
			}

			lMap = append(lMap, mmap)
		}

		ret = dynamodb.AttributeValue{
			L: lMap,
		}
	} else {
		ret = dynamodb.AttributeValue{
			BOOL: aws.Bool(false),
		}
	}

	return ret, nil
}

func convertDynamoToPermissions(perms *dynamodb.AttributeValue) (Permission, error) {
	ret := Permission{}
	for key, value := range perms.M {
		switch key {
		case "name":
			ret.Name = *value.S
		case "action":
			ret.Action = *value.S
		case "identifier":
			ret.Identifier = *value.S
		}
	}
	if ret.Name != "" {
		return ret, nil
	}

	return ret, fmt.Errorf("couldn't convert to permissions")
}

func convertDynamoPermsToPermissions(perms []*dynamodb.AttributeValue) ([]Permission, error) {
	ret := []Permission{}
	for _, perm := range perms {
		p, err := convertDynamoToPermissions(perm)
		if err != nil {
			return ret, fmt.Errorf("conveert dynamo perms to perms err: %w", err)
		}
		ret = append(ret, p)
	}

	return ret, nil
}

func convertDynamoToPermission(perm map[string]*dynamodb.AttributeValue) (Permissions, error) {
	ret := Permissions{}
	for key, value := range perm {
		switch key {
		case "permissions":
			perms, err := convertDynamoPermsToPermissions(value.L)
			if err != nil {
				return Permissions{}, fmt.Errorf("convert dynamo to perms: %w", err)
			}
			ret.Permissions = perms
		case "identifier":
			ret.Identifier = *value.S
		}
	}
	if ret.Identifier != "" {
		return ret, nil
	}

	return ret, fmt.Errorf("couldnt convert to permissions")
}
