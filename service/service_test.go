package service_test

import (
	"errors"
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
				Body:     `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541","permissions":[{"name":"carpark","action":"book"},{"name":"account","action":"view"},{"name":"account","action":"test","identifier":"*"},{"name":"account","action":"login"}]}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541","permissions":[{"name":"carpark","action":"book","identifier":"2298f676-8d7c-5e38-a04d-72b572f23541"},{"name":"account","action":"view","identifier":"2298f676-8d7c-5e38-a04d-72b572f23541"},{"name":"account","action":"test","identifier":"*"},{"name":"account","action":"login","identifier":"2298f676-8d7c-5e38-a04d-72b572f23541"}]}`,
			},
			err: nil,
		},
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/create",
				Body:     `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541","permissions":[{"name":"account","action":"view","identifier":"*"}]}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 0,
			},
			err: errors.New("permission identifier already exists: ConditionalCheckFailedException: The conditional request failed"),
		},

		// Update
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/update",
				Body:     `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541","permissions":[{"name":"account","action":"view","identifier":"*"}]}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541","permissions":[{"name":"carpark","action":"book","identifier":"2298f676-8d7c-5e38-a04d-72b572f23541"},{"name":"account","action":"view","identifier":"*"},{"name":"account","action":"test","identifier":"*"},{"name":"account","action":"login","identifier":"2298f676-8d7c-5e38-a04d-72b572f23541"}]}`,
			},
			err: nil,
		},

		// Retrieve
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/retrieve",
				Body:     `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541"}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541","permissions":[{"name":"carpark","action":"book","identifier":"2298f676-8d7c-5e38-a04d-72b572f23541"},{"name":"account","action":"view","identifier":"*"},{"name":"account","action":"test","identifier":"*"},{"name":"account","action":"login","identifier":"2298f676-8d7c-5e38-a04d-72b572f23541"}]}`,
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
				Body:     `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541","permissions":[{"name":"account","action":"login"}]}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541","status":"allowed"}`,
			},
			err: nil,
		},
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/allowed",
				Body:     `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541","permissions":[{"name":"carparks","action":"create"}]}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541","status":"denied"}`,
			},
			err: nil,
		},

		// Remove
		{
			request: events.APIGatewayProxyRequest{
				Resource: "/delete",
				Body:     `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541"}`,
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23541","status":"deleted"}`,
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
