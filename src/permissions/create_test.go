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
				Name: "tester1",
				Permission: "test1",
				User: "test1",
			},
			expect: permissions.Permission{
				ID: "dce68a10-cb43-5d69-a488-2280f3fd5eb3",
				Name: "tester1",
				User: "test1",
				Status: permissions.PermissionGood,
				AllowedTo: "test1",
				Company: true,
			},
			err:    nil,
		},
		{
			request: permissions.PermissionRequest{
				Name: "tester1",
				Permission: "test1",
				User: "test1",
			},
			expect: permissions.Permission{
				ID: "dce68a10-cb43-5d69-a488-2280f3fd5eb3",
				Name: "tester1",
				User: "test1",
				Status: permissions.PermissionBad,
				AllowedTo: "test1",
				Company: true,
			},
			err:    nil,
		},
	}

	for _, test := range tests {
		response, err := test.request.Create()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response)
	}

	tests[1].request.DeletePermission()
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
				ID: "ceebe569-be51-5636-b31c-935e5184cb26",
				Name: "tester",
				User: "test",
				Status: permissions.PermissionGood,
				AllowedTo: "test",
				Company: false,
			},
			err:    nil,
		},
		{
			request: permissions.PermissionRequest{
				Name: "tester",
				Permission: "test",
				User: "test",
			},
			expect: permissions.Permission{
				ID: "ceebe569-be51-5636-b31c-935e5184cb26",
				Name: "tester",
				User: "test",
				Status: permissions.PermissionBad,
				AllowedTo: "test",
				Company: false,
			},
			err:    nil,
		},
	}

	for _, test := range tests {
		response, err := test.request.CreateUser()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response)
	}
	tests[1].request.DeleteUser()
}