package permissions

import (
	"encoding/json"
	// "errors"
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"

	// "strings"
)

// PermissionResponse struct
type PermissionResponse struct {
	Permission Permission `json:"permission,omitempty"`
	Error      error      `json:"error,omitempty"`
}

// Permissions struct
type Permission struct {
	Identity    string        `json:"identity"`
	Permissions []Permissions `json:"permissions"`
}

// Permission struct
type Permissions struct {
	Name   string `json:"name"`
	Action string `json:"action"`
}

// PermissionRequest struct
type PermissionRequest struct {
	Identity    string
	Permissions []Permission
}

// PermissionRequestUpdate struct
type PermissionRequestUpdate struct {
	Old PermissionRequest
	New PermissionRequest
}

func (pr PermissionRequest) getUserUUID() string {
	u := uuid.NewV5(uuid.NamespaceDNS, fmt.Sprintf("https://permissions.carprk.com/user/%s", pr.Identity))
	return u.String()
}

func (pr PermissionRequest) getCompanyUUID() string {
	u := uuid.NewV5(uuid.NamespaceURL, fmt.Sprintf("https://permissions.carprk.com/company/%s", pr.Identity))

	return u.String()
}

// PermissionAstrix in-case we want a different wildcard
const PermissionAstrix = "*"

func ErrorResponse(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Println(fmt.Sprintf("err: %v", e))
	eErr := json.NewEncoder(w).Encode(PermissionResponse{
		Error: e,
	})
	if eErr != nil {
		fmt.Println(fmt.Sprintf("encode err: %v", eErr))
	}
}