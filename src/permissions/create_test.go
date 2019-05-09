package permissions_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"github.com/stretchr/testify/assert"
// 	"io/ioutil"
// 	"main/src/permissions"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )
//
// func TestCreate(t *testing.T) {
// 	tests := []struct{
// 		request permissions.PermissionRequest
// 		expect  permissions.Permission
// 		err     error
// 	}{
// 		{
// 			request: permissions.PermissionRequest{
// 				Name: "tester1",
// 				Permission: "test1",
// 				Identity: "test1",
// 			},
// 			expect: permissions.Permission{
// 				ID: "dce68a10-cb43-5d69-a488-2280f3fd5eb3",
// 				Name: "tester1",
// 				Identity: "test1",
// 				Status: permissions.PermissionGood,
// 				AllowedTo: "test1",
// 				Company: true,
// 			},
// 			err:    nil,
// 		},
// 		{
// 			request: permissions.PermissionRequest{
// 				Name: "tester1",
// 				Permission: "test1",
// 				Identity: "test1",
// 			},
// 			expect: permissions.Permission{
// 				ID: "dce68a10-cb43-5d69-a488-2280f3fd5eb3",
// 				Name: "tester1",
// 				Identity: "test1",
// 				Status: permissions.PermissionBad,
// 				AllowedTo: "test1",
// 				Company: true,
// 			},
// 			err:    nil,
// 		},
// 	}
//
// 	for _, test := range tests {
// 		response, err := test.request.Create()
// 		assert.IsType(t, test.err, err)
// 		assert.Equal(t, test.expect, response)
// 	}
//
// 	tests[0].request.DeletePermission()
// }
//
// func TestCreateUser(t *testing.T) {
// 	tests := []struct{
// 		request permissions.PermissionRequest
// 		expect  permissions.Permission
// 		err     error
// 	}{
// 		{
// 			request: permissions.PermissionRequest{
// 				Name: "tester",
// 				Permission: "test",
// 				Identity: "test",
// 			},
// 			expect: permissions.Permission{
// 				ID: "ceebe569-be51-5636-b31c-935e5184cb26",
// 				Name: "tester",
// 				Identity: "test",
// 				Status: permissions.PermissionGood,
// 				AllowedTo: "test",
// 				Company: false,
// 			},
// 			err:    nil,
// 		},
// 		{
// 			request: permissions.PermissionRequest{
// 				Name: "tester",
// 				Permission: "test",
// 				Identity: "test",
// 			},
// 			expect: permissions.Permission{
// 				ID: "ceebe569-be51-5636-b31c-935e5184cb26",
// 				Name: "tester",
// 				Identity: "test",
// 				Status: permissions.PermissionBad,
// 				AllowedTo: "test",
// 				Company: false,
// 			},
// 			err:    nil,
// 		},
// 	}
//
// 	for _, test := range tests {
// 		response, err := test.request.CreateUser()
// 		assert.IsType(t, test.err, err)
// 		assert.Equal(t, test.expect, response)
// 	}
// 	tests[1].request.DeleteUser()
// }
//
// func TestCreateUserHTTP(t *testing.T) {
// 	tests := []struct{
// 		request permissions.PermissionRequestHTTP
// 		expect permissions.PermissionResponse
// 		err error
// 	}{
// 		{
// 			request: permissions.PermissionRequestHTTP{
// 				Permission: "company:create",
// 				Identity: "keloran",
// 			},
// 			expect: permissions.PermissionResponse{
// 				Permission: permissions.Permission{
// 					ID: "ae85365b-b0c4-5a7b-83d0-152592c3f50d",
// 					Name: "company",
// 					AllowedTo: "create",
// 					Identity: "keloran",
// 					Status: permissions.PermissionGood,
// 					Company: false,
// 				},
// 			},
// 			err: nil,
// 		},
// 		{
// 			request: permissions.PermissionRequestHTTP{
// 				Permission: "company:create",
// 				Identity: "keloran",
// 			},
// 			expect: permissions.PermissionResponse{
// 				Permission: permissions.Permission{
// 					ID: "ae85365b-b0c4-5a7b-83d0-152592c3f50d",
// 					Name: "company",
// 					AllowedTo: "create",
// 					Identity: "keloran",
// 					Status: permissions.PermissionBad,
// 					Company: false,
// 				},
// 			},
// 			err: nil,
// 		},
// 	}
//
// 	for _, test := range tests {
// 		jpr, _ := json.Marshal(test.request)
// 		request, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jpr))
// 		response := httptest.NewRecorder()
// 		permissions.CreateUser(response, request)
// 		assert.Equal(t, 201, response.Code)
//
// 		body, _ := ioutil.ReadAll(response.Body)
// 		p := permissions.PermissionResponse{}
// 		json.Unmarshal(body, &p)
// 		assert.Equal(t, test.expect, p)
// 	}
//
// 	pr, _ := tests[0].request.ConvertToPermissionRequest()
// 	pr.DeleteUser()
// }
//
// func TestCreateHTTP(t *testing.T) {
// 	tests := []struct{
// 		request permissions.PermissionRequestHTTP
// 		expect permissions.PermissionResponse
// 		err error
// 	}{
// 		{
// 			request: permissions.PermissionRequestHTTP{
// 				Permission: "company:createUser",
// 				Identity: "carprkTest",
// 			},
// 			expect: permissions.PermissionResponse{
// 				Permission: permissions.Permission{
// 					ID: "589217cd-cb7b-5cab-b3fa-5fd924885962",
// 					Name: "company",
// 					AllowedTo: "createUser",
// 					Identity: "carprkTest",
// 					Status: permissions.PermissionGood,
// 					Company: true,
// 				},
// 			},
// 			err: nil,
// 		},
// 		{
// 			request: permissions.PermissionRequestHTTP{
// 				Permission: "company:createUser",
// 				Identity: "carprkTest",
// 			},
// 			expect: permissions.PermissionResponse{
// 				Permission: permissions.Permission{
// 					ID: "589217cd-cb7b-5cab-b3fa-5fd924885962",
// 					Name: "company",
// 					AllowedTo: "createUser",
// 					Identity: "carprkTest",
// 					Status: permissions.PermissionBad,
// 					Company: true,
// 				},
// 			},
// 			err: nil,
// 		},
// 	}
//
// 	for _, test := range tests {
// 		jpr, _ := json.Marshal(test.request)
// 		request, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jpr))
// 		response := httptest.NewRecorder()
// 		permissions.Create(response, request)
// 		assert.Equal(t, 201, response.Code)
//
// 		body, _ := ioutil.ReadAll(response.Body)
// 		p := permissions.PermissionResponse{}
// 		json.Unmarshal(body, &p)
// 		assert.Equal(t, test.expect, p)
// 	}
//
// 	pr, _ := tests[0].request.ConvertToPermissionRequest()
// 	pr.DeletePermission()
// }