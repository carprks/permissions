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

var serviceTests = []struct {
	name    string
	request events.APIGatewayProxyRequest
	expect  events.APIGatewayProxyResponse
	err     error
}{
	// Create
	{
		name: "create allowed",
		request: events.APIGatewayProxyRequest{
			Resource: "/create",
			Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"carpark","action":"book"},{"name":"account","action":"view"},{"name":"account","action":"test","identifier":"*"},{"name":"account","action":"login"}]}`,
		},
		expect: events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"carpark","action":"book","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"},{"name":"account","action":"view","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"},{"name":"account","action":"test","identifier":"*"},{"name":"account","action":"login","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"}]}`,
		},
	},
	{
		name: "create failed",
		request: events.APIGatewayProxyRequest{
			Resource: "/create",
			Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"account","action":"view","identifier":"*"}]}`,
		},
		expect: events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       `handler err: create entry err: permission identifier already exists: %!w(string=ConditionalCheckFailedException)`,
		},
	},

	// Update
	{
		name: "update success",
		request: events.APIGatewayProxyRequest{
			Resource: "/update",
			Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"account","action":"view","identifier":"*"}]}`,
		},
		expect: events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"carpark","action":"book","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"},{"name":"account","action":"view","identifier":"*"},{"name":"account","action":"test","identifier":"*"},{"name":"account","action":"login","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"}]}`,
		},
	},

	// Retrieve
	{
		name: "retrieve success",
		request: events.APIGatewayProxyRequest{
			Resource: "/retrieve",
			Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"}`,
		},
		expect: events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"carpark","action":"book","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"},{"name":"account","action":"view","identifier":"*"},{"name":"account","action":"test","identifier":"*"},{"name":"account","action":"login","identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"}]}`,
		},
	},
	{
		name: "retrieve failed",
		request: events.APIGatewayProxyRequest{
			Resource: "/retrieve",
			Body:     `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23542"}`,
		},
		expect: events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       `{"identifier":"2298f676-8d7c-5e38-a04d-72b572f23542","status":"no permissions"}`,
		},
	},

	// Check
	{
		name: "allowed success",
		request: events.APIGatewayProxyRequest{
			Resource: "/allowed",
			Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"account","action":"login"}]}`,
		},
		expect: events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","status":"allowed"}`,
		},
	},
	{
		name: "allowed failed",
		request: events.APIGatewayProxyRequest{
			Resource: "/allowed",
			Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"carparks","action":"create"}]}`,
		},
		expect: events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","status":"denied"}`,
		},
	},

	// Diety
	{
		name: "make deity",
		request: events.APIGatewayProxyRequest{
			Resource: "/update/deity",
			Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"}`,
		},
		expect: events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","permissions":[{"name":"*","action":"*","identifier":"*"}]}`,
		},
	},

	// Remove
	{
		name: "remove success",
		request: events.APIGatewayProxyRequest{
			Resource: "/delete",
			Body:     `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250"}`,
		},
		expect: events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       `{"identifier":"5f46cf19-5399-55e3-aa62-0e7c19382250","status":"deleted"}`,
		},
	},
}

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

	for _, test := range serviceTests {
		t.Run(test.name, func(t *testing.T) {
			response, err := service.Handler(test.request)
			passed := assert.IsType(t, test.err, err)
			if !passed {
				t.Errorf("service test err: %w, request: %v", err, test.request)
			}
			passed = assert.Equal(t, test.expect, response)
			if !passed {
				t.Errorf("service test not equal: %v", test.request)
			}
		})
	}
}

func BenchmarkHandler(b *testing.B) {
	b.ReportAllocs()

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

	b.ResetTimer()
	for _, test := range serviceTests {
		b.Run(test.name, func(b *testing.B) {
			b.StopTimer()

			_, err := service.Handler(test.request)
			passed := assert.IsType(b, test.err, err)
			if !passed {
				b.Errorf("service test err: %w, request: %v", err, test.request)
			}

			b.StartTimer()
		})
	}
}
