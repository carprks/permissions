package service_test

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/carprks/permissions/service"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	if len(os.Args) >= 1 {
		for _, env := range os.Args {
			if env == "localDev" {
				err := godotenv.Load()
				if err != nil {
					fmt.Println(fmt.Sprintf("godotenv err: %v", err))
				}
			}
		}
	}

	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  events.APIGatewayProxyResponse
		err     error
	}{
		// Create
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/create",
				Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"carpark","action":"book"},{"name":"account","action":"view"},{"name":"account","action":"test","identifier":"*"},{"name":"account","action":"login"}]}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"carpark","action":"book","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"},{"name":"account","action":"view","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"},{"name":"account","action":"test","identifier":"*"},{"name":"account","action":"login","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"}]}`,
			},
			err: nil,
		},
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/create",
				Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"account","action":"view","identifier":"*"}]}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 400,
				Body: `handler err: create entry err: permission identifier already exists: %!w(string=ConditionalCheckFailedException)`,
			},
		},

		// Update
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/update",
				Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"account","action":"view","identifier":"*"}]}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"carpark","action":"book","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"},{"name":"account","action":"view","identifier":"*"},{"name":"account","action":"test","identifier":"*"},{"name":"account","action":"login","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"}]}`,
			},
			err: nil,
		},

		// Retrieve
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/retrieve",
				Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"carpark","action":"book","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"},{"name":"account","action":"view","identifier":"*"},{"name":"account","action":"test","identifier":"*"},{"name":"account","action":"login","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"}]}`,
			},
			err: nil,
		},
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/retrieve",
				Body:     `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23542"}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23542","status":"no permissions"}`,
			},
			err: nil,
		},

		// Check
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/allowed",
				Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"account","action":"login"}]}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","status":"allowed"}`,
			},
			err: nil,
		},
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/allowed",
				Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"carparks","action":"create"}]}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","status":"denied"}`,
			},
			err: nil,
		},

		// Remove
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/delete",
				Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","status":"deleted"}`,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		response, err := service.Handler(test.request)
		passed := assert.IsType(t, test.err, err)
		if !passed {
			fmt.Println(fmt.Sprintf("service test err: %v, request: %v", err, test.request))
		}
		passed = assert.Equal(t, test.expect, response)
		if !passed {
			fmt.Println(fmt.Sprintf("service test not equal: %v", test.request))
		}
	}
}
