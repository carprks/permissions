package healthcheck_test

import (
	"github.com/stretchr/testify/assert"
	"main/src/healthcheck"
	"testing"
)

func TestHealthCheck_Check(t *testing.T) {
	tests := []struct{
		request healthcheck.HealthCheck
		response healthcheck.Health
		err error
	}{
		{
			request: healthcheck.HealthCheck{
				Name: "test1",
				URL: "test1.com",
				Dependencies: "",
			},
			response: healthcheck.Health{
				Name: "test1",
				URL: "test1.com",
				Status: healthcheck.HealthPass,
				Dependencies: nil,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		response, err := test.request.Check()
		assert.Equal(t, test.err, err)
		assert.Equal(t, test.response, response)
	}
}
