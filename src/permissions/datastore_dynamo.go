package permissions

import (
    "errors"
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "os"
)

func stringMapToDynamoMap(perms map[string]string) map[string]*dynamodb.AttributeValue {
    ret := map[string]*dynamodb.AttributeValue{}

    if len(perms) >= 1 {
        for key, value := range perms {
            ret[key] = &dynamodb.AttributeValue{
                S: aws.String(value),
            }
        }
    }

    return ret
}

func (p Permissions) storeDynamo() (Permissions, error) {
    s, err := session.NewSession(&aws.Config{
        Region: aws.String(os.Getenv("AWS_DB_REGION")),
        Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
    })
    if err != nil {
        fmt.Println(fmt.Sprintf("Store Dynamo Error: %v", err))
        return p, err
    }
    svc := dynamodb.New(s)
    input := &dynamodb.PutItemInput{
        Item: map[string]*dynamodb.AttributeValue{
            "identifier": {
                S: aws.String(p.ID),
            },
            "permissions": {
                M: stringMapToDynamoMap(p.MapPermissions()),
            },
            "identity": {
                S: aws.String(p.Identity),
            },
            "status": {
                S: aws.String(string(p.Status)),
            },
            "company": {
                BOOL: aws.Bool(p.Company),
            },
        },
        TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
    }
    _, err = svc.PutItem(input)
    if err != nil {
        fmt.Println(fmt.Sprintf("Dynamo Put Item: %v", err))
        return p, err
    }

    return p, nil
}

func (p Permissions) updateDynamo() (Permissions, error) {
    s, err := session.NewSession(&aws.Config{
        Region: aws.String(os.Getenv("AWS_DB_REGION")),
        Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
    })
    if err != nil {
        fmt.Println(fmt.Sprintf("Update Dynamo Error: %v", err))
        return Permissions{}, err
    }
    svc := dynamodb.New(s)
    input := &dynamodb.UpdateItemInput{
        ExpressionAttributeNames: map[string]*string{
            "#PERMISSION": aws.String("permission"),
            "#IDENTITY": aws.String("identity"),
            "#STATUS": aws.String("status"),
            "#COMPANY": aws.String("company"),
        },
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":permission": {
                M: stringMapToDynamoMap(p.MapPermissions()),
            },
            ":identity": {
                S: aws.String(p.Identity),
            },
            ":status": {
                S: aws.String(string(p.Status)),
            },
            ":company": {
                BOOL: aws.Bool(p.Company),
            },
        },
        Key: map[string]*dynamodb.AttributeValue{
            "identifier": {
                S: aws.String(p.ID),
            },
        },
        ReturnValues: aws.String("ALL_NEW"),
        TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
        UpdateExpression: aws.String("SET #PERMISSION = :permission, #IDENTITY = :identity, #STATUS = :status, #COMPANY = :company"),
    }
    _, err = svc.UpdateItem(input)
    if err != nil {
        fmt.Println(fmt.Sprintf("Update Dynamo Input Error: %v", err))
        return Permissions{}, err
    }
    return p, nil
}

func (p Permissions) deleteDynamo() (Permissions, error) {
    s, err := session.NewSession(&aws.Config{
        Region: aws.String(os.Getenv("AWS_DB_REGION")),
        Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
    })
    if err != nil {
        fmt.Println(fmt.Sprintf("Delete Dynamo Error: %v", err))
        return Permissions{}, err
    }
    svc := dynamodb.New(s)
    input := &dynamodb.DeleteItemInput{
        Key: map[string]*dynamodb.AttributeValue{
            "identifier": {
                S: aws.String(p.ID),
            },
        },
        TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
    }
    _, err = svc.DeleteItem(input)
    if err != nil {
        fmt.Println(fmt.Sprintf("Delete Dynamo Action Error: %v", err))
        return Permissions{}, err
    }
    return p, nil
}

func (p Permissions) retrieveAstrixDynamo() (Permissions, error) {
    s, err := session.NewSession(&aws.Config{
        Region: aws.String(os.Getenv("AWS_DB_REGION")),
        Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
    })
    if err != nil {
        fmt.Println(fmt.Sprintf("Retrieve Astrix Dynamo Error: %v", err))
        return Permissions{}, err
    }
    svc := dynamodb.New(s)
    input := &dynamodb.ScanInput{
        TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
        ExpressionAttributeNames: map[string]*string{
            "#IDENTITY": aws.String("identity"),
        },
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":identity": {
                S: aws.String(p.Identity),
            },
        },
        FilterExpression: aws.String("#IDENTITY = :identity AND #NAME = :name"),
    }
    result, err := svc.Scan(input)
    if err != nil {
        fmt.Println(fmt.Sprintf("Retrieve Astrix Dynamo Scan: %v", err))
        return Permissions{}, err
    }

    if len(result.Items) >= 1 {
        pr := Permissions{}

        for i := 0; i < len(result.Items); i++ {
            perm := result.Items[i]

            // // find specific
            // if *perm["permission"].S == p.Permissions {
            //     pr = Permissions{
            //         ID: *perm["identifier"].S,
            //         Identity: *perm["identity"].S,
            //         Permissions: *perm["permission"].S,
            //         Name: *perm["name"].S,
            //         Company: *perm["company"].BOOL,
            //         Status: getStatus(*perm["status"].S),
            //     }
            // }

            // permission is an astrix
            if *perm["permission"].S == PermissionAstrix {
                pr = Permissions{
                    ID: *perm["identifier"].S,
                    Identity: *perm["identity"].S,
                    // Permissions: *perm["permission"].S,
                    Company: *perm["company"].BOOL,
                    Status: getStatus(*perm["status"].S),
                }
            }
        }

        return pr, nil
    }

    return Permission{}, errors.New("no permission entry")
}

func (p Permissions) retrieveDynamo() (Permissions, error) {
    s, err := session.NewSession(&aws.Config{
        Region: aws.String(os.Getenv("AWS_DB_REGION")),
        Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
    })
    if err != nil {
        fmt.Println(fmt.Sprintf("Retrive Dynamo Error: %v", err))
        return Permissions{}, err
    }
    svc := dynamodb.New(s)
    input := &dynamodb.GetItemInput{
        Key: map[string]*dynamodb.AttributeValue{
            "identifier": {
                S: aws.String(p.ID),
            },
        },
        TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
    }
    result, err := svc.GetItem(input)
    if err != nil {
        fmt.Println(fmt.Sprintf("Dynamo Get Item: %v", err))
        return Permissions{}, err
    }
    if len(result.Item) >= 1 {
        pp := Permissions{
            ID: *result.Item["identifier"].S,
            Name: *result.Item["name"].S,
            Status: getStatus(*result.Item["status"].S),
            Identity: *result.Item["identity"].S,
            Permissions: *result.Item["permissions"].M,
            Company: *result.Item["company"].BOOL,
        }

        return pp, nil
    }
    return Permissions{}, errors.New("no permission entry")
}

func (p Permissions) retrieveAllDynamo() ([]Permissions, error) {
    pr := []Permissions{}

    s, err := session.NewSession(&aws.Config{
        Region: aws.String(os.Getenv("AWS_DB_REGION")),
        Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
    })
    if err != nil {
        fmt.Println(fmt.Sprintf("Retrive Dynamo Error: %v", err))
        return pr, err
    }
    svc := dynamodb.New(s)
    input := &dynamodb.ScanInput{
        TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
    }
    result, err := svc.Scan(input)
    if err != nil {
        fmt.Println(fmt.Sprintf("Dynamo Get Item: %v", err))
        return []Permissions{}, err
    }
    if len(result.Items) >= 1 {
        for i := 0; i < len(result.Items); i++ {
            item := result.Items[i]

            pp := Permissions{
                ID: *item["identifier"].S,
                Name: *item["name"].S,
                Status: getStatus(*item["status"].S),
                Identity: *item["identity"].S,
                Permissions: *item["permission"].S,
                Company: *item["company"].BOOL,
            }
            pr = append(pr, pp)
        }
        return pr, nil
    }

    return []Permissions{}, nil
}

// DeleteTable remove the whole table
func DeleteTable() error {
    tableExists := tableExists()
    if !tableExists {
        s, err := session.NewSession(&aws.Config{
            Region:   aws.String(os.Getenv("AWS_DB_REGION")),
            Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
        })
        if err != nil {
            fmt.Println(fmt.Sprintf("Delete Table Session: %v", err))
            return err
        }
        svc := dynamodb.New(s)
        _, err = svc.DeleteTable(&dynamodb.DeleteTableInput{
            TableName: aws.String("AWS_DB_TABLE"),
        })
        if err != nil {
            fmt.Println(fmt.Sprintf("Delete Table: %v", err))
            return err
        }
    }
    return nil
}

// CreateTable create a new table
func CreateTable() error {
    tableExists := tableExists()
    if !tableExists {
        s, err := session.NewSession(&aws.Config{
            Region:   aws.String(os.Getenv("AWS_DB_REGION")),
            Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
        })
        if err != nil {
            fmt.Println(fmt.Sprintf("Create Table Session: %v", err))
            return err
        }
        svc := dynamodb.New(s)
        _, err = svc.CreateTable(&dynamodb.CreateTableInput{
            AttributeDefinitions: []*dynamodb.AttributeDefinition{
                {
                    AttributeName: aws.String("identifier"),
                    AttributeType: aws.String("S"),
                },
            },
            KeySchema: []*dynamodb.KeySchemaElement{
                {
                    AttributeName: aws.String("identifier"),
                    KeyType:       aws.String("HASH"),
                },
            },
            ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
                ReadCapacityUnits:  aws.Int64(5),
                WriteCapacityUnits: aws.Int64(5),
            },
            TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
        })
        if err != nil {
            fmt.Println(fmt.Sprintf("Create Table: %v", err))
            return err
        }
    }
    return nil
}

func tableExists() bool {
    s, err := session.NewSession(&aws.Config{
        Region: aws.String(os.Getenv("AWS_DB_REGION")),
        Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
    })
    if err != nil {
        fmt.Println(fmt.Sprintf("Table Exists Session: %v", err))
        return false
    }
    svc := dynamodb.New(s)
    out, err := svc.ListTables(&dynamodb.ListTablesInput{
        // ExclusiveStartTableName: aws.String(os.Getenv("AWS_DB_TABLE")),
    })
    if err != nil {
        fmt.Println(fmt.Sprintf("Table Exists List: %v", err))
        return false
    }
    if len(out.TableNames) >= 1 {
        for i := 0; i < len(out.TableNames);i ++ {
            name := out.TableNames[i]
            if *name == os.Getenv("AWS_DB_TABLE") {
                return true
            }
        }
    }
    return false
}