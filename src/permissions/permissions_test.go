package permissions_test

import (
	"github.com/stretchr/testify/assert"
	"main/src/permissions"
	"testing"
)

func TestCreate(t *testing.T) {
	pr := permissions.PermissionRequest{}

	tests := []struct {
		request permissions.Permission
		expect  string
		err     error
	}{
		{
			expect: permissions.Permission{},
			err:    nil,
		},
	}

	for _, test := range tests {
		response, err := pr.Create()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response)
	}
}
