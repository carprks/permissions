package permissions_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"main/src/permissions"
	"testing"
)

const testRetrieve = "testRetrieve"
const testRetrieveFake = "testRetrieveFake"
const testAll = "testAll"

func TestRetrievePermissions(t *testing.T) {
	tests := []struct{
		request     permissions.PermissionRequest
		expect      permissions.Permission
		err         error
		skipCreate  bool
	}{
		{
			request: permissions.PermissionRequest{
				Name: testRetrieve,
				Permission: testRetrieve,
				User: testRetrieve,
			},
			expect: permissions.Permission{
				ID: "73b780c8-b32b-5758-9c0c-069f38899045",
				Name: testRetrieve,
				AllowedTo: testRetrieve,
				User: testRetrieve,
				Company: true,
				Status: permissions.PermissionGood,
			},
			err: nil,
			skipCreate: false,
		},{
			request: permissions.PermissionRequest{
				Name: testRetrieveFake,
				Permission: testRetrieveFake,
				User: testRetrieveFake,
			},
			expect: permissions.Permission{
				ID: "abb022fb-ee38-5ac7-b263-428517f522f7",
				Name: testRetrieveFake,
				AllowedTo: testRetrieveFake,
				User: testRetrieveFake,
				Status: permissions.PermissionBad,
				Company: true,
			},
			err: errors.New("no permission entry"),
			skipCreate: true,
		},
	}

	for _, test := range tests {
		if !test.skipCreate {
			test.request.Create()
		}

		response, err := test.request.RetrievePermissions()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response)

		test.request.DeletePermission()
	}
}

func TestRetrieveUser(t *testing.T) {
	tests := []struct{
		request permissions.PermissionRequest
		expect permissions.Permission
		err error
		skipCreate bool
	}{
		{
			request: permissions.PermissionRequest{
				Name: testRetrieve,
				Permission: testRetrieve,
				User: testRetrieve,
			},
			expect: permissions.Permission{
				ID: "cb14809b-d814-52d4-8cde-336c592662bc",
				Name: testRetrieve,
				AllowedTo: testRetrieve,
				User: testRetrieve,
				Status: permissions.PermissionGood,
				Company: false,
			},
			err: nil,
			skipCreate: false,
		},
		{
			request: permissions.PermissionRequest{
				Name: testRetrieveFake,
				Permission: testRetrieveFake,
				User: testRetrieveFake,
			},
			expect: permissions.Permission{
				ID: "7a4bc518-32a6-528a-b695-f45c7d20fb1c",
				Name: testRetrieveFake,
				User: testRetrieveFake,
				AllowedTo: testRetrieveFake,
				Status: permissions.PermissionBad,
				Company: false,
			},
			err: errors.New("no permission entry"),
			skipCreate: true,
		},
	}

	for _, test := range tests {
		if !test.skipCreate {
			test.request.CreateUser()
		}

		response, err := test.request.RetrieveUser()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response)

		test.request.DeleteUser()
	}
}

func TestRetrievePermissionAstrix(t *testing.T) {
	// Test all permission
	pc := permissions.PermissionRequest{
		Name: testAll,
		Permission: "*",
		User: testAll,
	}
	pc.Create()

	pr := permissions.PermissionRequest{
		Name: testAll,
		Permission: "bob",
		User: testAll,
	}
	response, err := pr.RetrievePermissions()
	assert.IsType(t, nil, err)
	assert.Equal(t, permissions.Permission{
		ID: "b2e0917e-d18d-5f00-bbca-bec305dc0302",
		Name: testAll,
		User: testAll,
		AllowedTo: "*",
		Company: true,
		Status: permissions.PermissionGood,
	}, response)

	pc.DeleteUser()
}

