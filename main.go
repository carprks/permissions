package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/carprks/permissions/service"
)

func main() {
	lambda.Start(service.Handler)
}
