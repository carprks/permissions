package permissions_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/carprks/permissions/src/permissions"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	if os.Getenv("AWS_DB_TABLE") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(fmt.Sprintf("godotenv err: %v", err))
		}
	}

	tests := []struct{
		request permissions.Permissions
		expect permissions.PermissionResponse
		err error
	}{
		{
			request: permissions.Permissions{
				Identity: "testerHTTP",
				Permissions: []permissions.Permission{
					{
						Action: "create",
						Name: "account",
					},
				},
			},
			expect: permissions.PermissionResponse{
				Permissions: permissions.Permissions{
					Identity: "testerHTTP",
					Permissions: []permissions.Permission{
						{
							Action: "create",
							Name:   "account",
							Identifier: "testerHTTP",
						},
					},
				},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		jpr, _ := json.Marshal(test.request)
		request, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jpr))
		response := httptest.NewRecorder()
		permissions.Create(response, request)
		assert.Equal(t, 201, response.Code)

		body, _ := ioutil.ReadAll(response.Body)
		// fmt.Println(fmt.Sprintf("Body: %s", string(body)))
		p := permissions.PermissionResponse{}
		err := json.Unmarshal(body, &p)
		correct := assert.IsType(t, test.err, err)
		if !correct {
			fmt.Println(fmt.Sprintf("test create err: %v", err))
		}
		assert.Equal(t, test.expect, p)
	}

	for _, test := range tests {
		_, err := test.request.DeleteEntry()
		if err != nil {
			fmt.Println(fmt.Sprintf("create test delete err: %v", err))
		}
	}
}
