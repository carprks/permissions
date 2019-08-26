package permissions_test

import (
	"fmt"
	"github.com/carprks/permissions/src/permissions"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPermission_CreateEntry(t *testing.T) {
	if os.Getenv("AWS_DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	perm := permissions.Permissions{
		Identity: "tester",
		Permissions: []permissions.Permission{
			{
				Name: "account",
				Action: "create",
			},
		},
	}
	tests := []struct{
		request permissions.Permissions
		expect permissions.Permissions
		err error
	}{
		{
			request: perm,
			expect: permissions.Permissions{
        Identity:    "tester",
        Permissions: []permissions.Permission{
          {
            Name: "account",
            Action: "create",
            Identifier: "tester",
          },
        },
      },
			err: nil,
		},
	}

	for _, test := range tests {
		response, err := perm.CreateEntry()
		correct := assert.IsType(t, test.err, err)
		if !correct {
			fmt.Println(fmt.Sprintf("create test err: %v", err))
		}
		assert.Equal(t, test.expect, response)
	}
}

func TestPermission_UpdateEntry(t *testing.T) {
	if os.Getenv("AWS_DB_TABLE") != "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	orig := permissions.Permissions{
		Identity: "tester",
		Permissions: []permissions.Permission{
			{
				Name: "account",
				Action: "create",
			},
		},
	}
	n := permissions.Permissions{
		Identity: "tester",
		Permissions: []permissions.Permission{
			{
				Name: "account",
				Action: "create",
				Identifier: "tester",
			},
			{
				Name: "*",
				Action: "*",
				Identifier: "*",
			},
		},
	}

	tests := []struct{
		request permissions.Permissions
		update permissions.Permissions
		expect permissions.Permissions
		err error
	}{
		{
			request: orig,
			update: n,
			expect: n,
			err: nil,
		},
	}

	for _, test := range tests {
		response, err := test.request.UpdateEntry(test.update)
		correct := assert.IsType(t, test.err, err)
		if !correct {
			fmt.Println(fmt.Sprintf("update test err: %v", err))
		}
		assert.Equal(t, test.expect, response)
	}
}

func TestScanEntries(t *testing.T) {
	if os.Getenv("AWS_DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	tests := []struct{
		expect int
		err error
	}{
		{
			expect: 1,
			err: nil,
		},
	}

	for _, test := range tests {
		response, err := permissions.ScanEntries()
		assert.IsType(t, test.err, err)
		assert.GreaterOrEqual(t, len(response), test.expect)
	}
}

func TestPermission_RetrieveEntry(t *testing.T) {
	if os.Getenv("AWS_DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	perm := permissions.Permissions{
		Identity: "tester",
		Permissions: []permissions.Permission{
			{
				Name: "account",
				Action: "create",
				Identifier: "tester",
			},
			{
				Name: "*",
				Action: "*",
				Identifier: "*",
			},
		},
	}

	tests := []struct{
		request permissions.Permissions
		expect permissions.Permissions
		err error
	}{
		{
			request: perm,
			expect: perm,
			err: nil,
		},
	}

	for _, test := range tests {
		resp, err := test.request.RetrieveEntry()
		correct := assert.IsType(t, test.err, err)
		if !correct {
			fmt.Println(fmt.Sprintf("retrieve test err: %v", err))
		}
		assert.Equal(t, test.expect, resp)
	}
}

func TestPermission_DeleteEntry(t *testing.T) {
	if os.Getenv("AWS_DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	perm := permissions.Permissions{
		Identity: "tester",
	}

	tests := []struct{
		request permissions.Permissions
		expect permissions.Permissions
		err error
	}{
		{
			request: perm,
			expect: permissions.Permissions{},
			err: nil,
		},
	}

	for _, test := range tests {
		response, err := test.request.DeleteEntry()
		correct := assert.IsType(t, test.err, err)
		if !correct {
			fmt.Println(fmt.Sprintf("retrieve test err: %v", err))
		}
		assert.Equal(t, test.expect, response)
	}
}
