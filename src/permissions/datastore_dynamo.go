package permissions

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "os"
)

func (p Permission) storeDynamo() (Permission, error) {
    s, err := session.NewSession(&aws.Config{
        Region: aws.String(os.Getenv("AWS_DB_REGION")),
        Endpoint: aws.String(os.Getenv("AWS_DB_ENDPOINT")),
    })
    if err != nil {
        fmt.Println(fmt.Sprintf("Store Dynamo Error: %v", err))
    }

    svc := dynamodb.New(s)
    input := &dynamodb.PutItemInput{
        Item: map[string]*dynamodb.AttributeValue{
            "identifier": {
                S: aws.String(p.ID),
            },
            "permission": {
                S: aws.String(p.AllowedTo),
            },
            "name": {
                S: aws.String(p.Name),
            },
            "user": {
                S: aws.String(p.User),
            },
            "status": {
                S: aws.String(string(p.Status)),
            },
        },
        TableName: aws.String(os.Getenv("AWS_DB_TABLE")),
    }

    _, err = svc.PutItem(input)
    if err != nil {
        fmt.Println(fmt.Sprintf("Dynamo Put Item: %v", err))
    }

    return p, nil
}

func (p Permission) updateDynamo() (Permission, error) {
    return Permission{}, nil
}

func (p Permission) deleteDynamo() (Permission, error) {
    return Permission{}, nil
}

func (p Permission) retrieveDynamo() (Permission, error) {
    return Permission{}, nil
}
