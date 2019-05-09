package permissions_test

import (
	"github.com/stretchr/testify/assert"
	"main/src/permissions"
	"testing"
)

const testCreate = "testCreate"
const testUpdate = "testUpdate"

func TestUpdatePermission(t *testing.T) {
	tests := []struct{
		create permissions.PermissionRequest
		update permissions.PermissionRequestUpdate
		expect permissions.Permission
		err error
	}{
		{
			create: permissions.PermissionRequest{
				Name: testCreate,
				Permission: testCreate,
				Identity: testCreate,
			},
			update: permissions.PermissionRequestUpdate{
				Old: permissions.PermissionRequest{
					Name: testCreate,
					Permission: testCreate,
					Identity: testCreate,
				},
				New: permissions.PermissionRequest{
					Name: testCreate,
					Permission: testUpdate,
					Identity: testCreate,
				},
			},
			expect: permissions.Permission{
				ID: "57186896-4148-5e58-9712-04ed2ba25a77",
				Identity: testCreate,
				AllowedTo: testUpdate,
				Name: testCreate,
				Company: true,
				Status: permissions.PermissionGood,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		test.create.Create()
		resp, err := test.update.UpdatePermission()

		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, resp)
	}

	tests[0].update.New.DeletePermission()
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		create permissions.PermissionRequest
		update permissions.PermissionRequestUpdate
		expect permissions.Permission
		err error
	}{
		{
			create: permissions.PermissionRequest{
				Name: testCreate,
				Permission: testCreate,
				Identity: testCreate,
			},
			update: permissions.PermissionRequestUpdate{
				Old: permissions.PermissionRequest{
					Name: testCreate,
					Permission: testCreate,
					Identity: testCreate,
				},
				New: permissions.PermissionRequest{
					Name: testCreate,
					Permission: testUpdate,
					Identity: testCreate,
				},
			},
			expect: permissions.Permission{
				ID: "0c044928-cdd8-50f6-b1c7-ff1abf057fbe",
				Identity: testCreate,
				AllowedTo: testUpdate,
				Name: testCreate,
				Company: false,
				Status: permissions.PermissionGood,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		test.create.CreateUser()

		resp, err := test.update.UpdateUser()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, resp)
	}

	tests[0].update.New.DeleteUser()
}