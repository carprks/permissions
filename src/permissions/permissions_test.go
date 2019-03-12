package permissions_test

import (
	"github.com/stretchr/testify/assert"
	"main/src/permissions"
	"testing"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		request permissions.Permission
		expect  permissions.Permission
		err     error
	}{
		{
			expect: permissions.Permission{
				Name: "tester",
			},
			err:    nil,
		},
	}

	for _, test := range tests {
		pr := permissions.PermissionRequest{
			Name: "tester",
		}

		response, err := pr.Create()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect.Name, response.Name)
	}
}
