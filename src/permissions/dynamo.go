package permissions

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

// CreateEntry create the permissions
func (p Permission)CreateEntry() (Permission, error) {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DB_REGION")),
		Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
	})
	if err != nil {
		return Permission{}, err
	}
	perms, err := convertPermissionsToDynamo(p.Permissions)
	if err != nil {
		return Permission{}, err
	}
	svc := dynamodb.New(s)
	input := &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
		Item: map[string]*dynamodb.AttributeValue{
			"identity": {
				S: aws.String(p.Identity),
			},
			"permissions": &perms,
		},
		ConditionExpression: aws.String("attribute_not_exists(#IDENTITY)"),
		ExpressionAttributeNames: map[string]*string{
			"#IDENTITY": aws.String("identity"),
		},
	}
	_, putErr := svc.PutItem(input)
	if putErr != nil {
		if awsErr, ok := putErr.(awserr.Error); ok {
			switch awsErr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				return Permission{}, fmt.Errorf("permission identity already exists: %v", awsErr)
			case "ValidationException":
				fmt.Println(fmt.Sprintf("validation err reason: %v", input))
				return Permission{}, fmt.Errorf("validation error: %v", awsErr)
			default:
				fmt.Println(fmt.Sprintf("unknown code err reason: %v", input))
				return Permission{}, fmt.Errorf("unknown code err: %v", awsErr)
			}
		} else {
			return Permission{}, fmt.Errorf("unknown err: %v", putErr)
		}
	}

	return p, nil
}

// RetrieveEntry get the permissions
func (p Permission)RetrieveEntry() (Permission, error) {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DB_REGION")),
		Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
	})
	if err != nil {
		return Permission{}, err
	}
	svc := dynamodb.New(s)
	input := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"identity": {
				S: aws.String(p.Identity),
			},
		},
	}
	result, err := svc.GetItem(input)
	if err != nil {
		return Permission{}, err
	}

	return convertDynamoToPermission(result.Item)
}

// UpdateEntry alter the permissions
func (p Permission)UpdateEntry(n Permission) (Permission, error) {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DB_REGION")),
		Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
	})
	if err != nil {
		return Permission{}, err
	}
	perms, err := convertPermissionsToDynamo(n.Permissions)
	if err != nil {
		return Permission{}, err
	}
	svc := dynamodb.New(s)
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
		ExpressionAttributeNames: map[string]*string{
			"#PERMISSIONS": aws.String("permissions"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":permissions": &perms,
		},
		Key: map[string]*dynamodb.AttributeValue{
			"identity": {
				S: aws.String(p.Identity),
			},
		},
		ReturnValues: aws.String("ALL_NEW"),
		UpdateExpression: aws.String("SET #PERMISSIONS = :permissions"),
	}
	ret, err := svc.UpdateItem(input)
	if err != nil {
		return Permission{}, err
	}

	return convertDynamoToPermission(ret.Attributes)
}

// DeleteEntry remove the permissions
func (p Permission)DeleteEntry() (Permission, error) {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DB_REGION")),
		Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
	})
	if err != nil {
		return Permission{}, err
	}
	svc := dynamodb.New(s)
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"identity": {
				S: aws.String(p.Identity),
			},
		},
	}
	_, err = svc.DeleteItem(input)
	if err != nil {
		return Permission{}, err
	}

	return Permission{}, nil
}

// ScanEntries get all the permisisons
func ScanEntries() ([]Permission, error) {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DB_REGION")),
		Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
	})
	if err != nil {
		return []Permission{}, err
	}
	svc := dynamodb.New(s)
	input := &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
	}
	result, err := svc.Scan(input)
	if err != nil {
		return []Permission{}, err
	}
	itemLen := len(result.Items)
	if itemLen >= 1 {
		perms := []Permission{}
		for i := 0; i < itemLen; i++ {
			item := result.Items[i]
			perm, err := convertDynamoToPermission(item)
			if err != nil {
				return perms, fmt.Errorf("couldnt convert perm: %v", err)
			}
			perms = append(perms, perm)
		}

		return perms, nil
	}

	return []Permission{}, nil
}

func convertPermissionsToDynamo(perms []Permissions) (dynamodb.AttributeValue, error) {
	ret := dynamodb.AttributeValue{}
	lMap := []*dynamodb.AttributeValue{}

	if len(perms) >= 1 {
		for _, perm := range perms {
			retMap := map[string]*dynamodb.AttributeValue{}
			retMap["name"] = &dynamodb.AttributeValue{
				S: aws.String(perm.Name),
			}
			retMap["action"] = &dynamodb.AttributeValue{
				S: aws.String(perm.Action),
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

func convertDynamoToPermissions(perms *dynamodb.AttributeValue) (Permissions, error) {
	ret := Permissions{}
	for key, value := range perms.M {
		switch key {
		case "name":
			ret.Name = *value.S
		case "action":
			ret.Action = *value.S
		}
	}
	if ret.Name != "" {
		return ret, nil
	}

	return ret, fmt.Errorf("couldnt convert to permissions")
}

func convertDynamoPermsToPermissions(perms []*dynamodb.AttributeValue) ([]Permissions, error) {
	ret := []Permissions{}
	for _, perm := range perms {
		p, err := convertDynamoToPermissions(perm)
		if err != nil {
			return ret, err
		}
		ret = append(ret, p)
	}

	return ret, nil
}

func convertDynamoToPermission(perm map[string]*dynamodb.AttributeValue) (Permission, error) {
	ret := Permission{}
	for key, value := range perm {
		switch key {
		case "permissions":
			perms, err := convertDynamoPermsToPermissions(value.L)
			if err != nil {
				return Permission{}, err
			}
			ret.Permissions = perms
		case "identity":
			ret.Identity = *value.S
		}
	}
	if ret.Identity != "" {
		return ret, nil
	}

	return ret, fmt.Errorf("couldnt convert to permission")
}