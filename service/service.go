package service

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

// Permissions struct
type Permissions struct {
	Identifier  string       `json:"identifier"`
	Permissions []Permission `json:"permissions,omitempty"`
	Status      string       `json:"status,omitempty"`
}

// Permission struct
type Permission struct {
	Name       string `json:"name"`
	Action     string `json:"action"`
	Identifier string `json:"identifier"`
}

func rest() (string, error) {
	return "", nil
}

// Handler ...
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp, err := rest()

	switch request.Resource {
	case "/create":
		resp, err = create(request.Body)
	case "/update":
		resp, err = update(request.Body)
	case "/delete":
		resp, err = delete(request.Body)
	case "/allowed":
		resp, err = allowed(request.Body)
	case "/retrieve":
		resp, err = retrieve(request.Body)
	case "/update/deity":
		resp, err = deity(request.Body)
	}

	if err != nil {
		fmt.Println(fmt.Sprintf("%v Err: %v", request.Resource, err))
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       fmt.Sprintf("handler err: %v", err),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       resp,
	}, nil
}
