package permissions_test

import (
	"github.com/stretchr/testify/assert"
	"main/src/permissions"
	"testing"
)

func TestCreate(t *testing.T) {
	tests := []struct{
		request permissions.PermissionRequest
		expect  permissions.Permission
		err     error
	}{
		{
			request: permissions.PermissionRequest{
				Name: "tester",
				Permission: "test",
				User: "test",
			},
			expect: permissions.Permission{
				ID: "388f43c8-4c32-580c-b192-c9ad602a01c2",
				Name: "tester",
				User: "test",
				Status: permissions.PermissionBad,
				AllowedTo: "test",
			},
			err:    nil,
		},
	}

	for _, test := range tests {
		response, err := test.request.Create()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response)
	}
}

func TestCreateUser(t *testing.T) {
	tests := []struct{
		request permissions.PermissionRequest
		expect  permissions.Permission
		err     error
	}{
		{
			request: permissions.PermissionRequest{
				Name: "tester",
				Permission: "test",
				User: "test",
			},
			expect: permissions.Permission{
				ID: "4ebc7ea4-7d5b-5e8d-9ade-9d13ada8b350",
				Name: "tester",
				User: "test",
				Status: permissions.PermissionBad,
				AllowedTo: "test",
			},
			err:    nil,
		},
	}

	for _, test := range tests {
		response, err := test.request.CreateUser()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response)
	}
}