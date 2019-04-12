package permissions

import (
	"fmt"
	"github.com/satori/go.uuid"
)

// Permission struct
type Permission struct {
	ID        string
	Name      string
	AllowedTo string
	User      string

	Status    PermissionStatus

	Company   bool
}

// PermissionRequest struct
type PermissionRequest struct {
	Name       string `json:"name"`
	Permission string `json:"permission"`
	User       string `json:"user"`
}

// PermissionRequestUpdate struct
type PermissionRequestUpdate struct {
	Old PermissionRequest `json:"oldPermission"`
	New PermissionRequest `json:"newPermission"`
}

// PermissionStatus string
type PermissionStatus string

const (
	// PermissionBad status of the permission
	PermissionBad  PermissionStatus = "Bad"

	// PermissionGood status of the permission
	PermissionGood PermissionStatus = "Good"
)

func getStatus(s string) PermissionStatus {
	if s == string(PermissionGood) {
		return PermissionGood
	}

	return PermissionBad
}

func (pr PermissionRequest) getUserUUID() string {
	u := uuid.NewV5(uuid.NamespaceDNS, fmt.Sprintf("https://permissions.carprk.com/user/%s/%s:%s", pr.User, pr.Name, pr.Permission))
	return u.String()
}

func (pr PermissionRequest) getCompanyUUID() string {
	u := uuid.NewV5(uuid.NamespaceURL, fmt.Sprintf("https://permissions.carprk.com/company/%s/%s:%s", pr.User, pr.Name, pr.Permission))

	return u.String()
}

// PermissionAstrix in-case we want a different wildcard
const PermissionAstrix = "*"

// checkExists
func (p Permission) checkExists() (bool, error) {
	resp, err := p.retrieveDynamo()
	if err != nil {
		if err.Error() == "no permission entry" {
			return false, nil
		}

		return false, err
	}

	// it has an id so it must exist
	if resp.ID != "" {
		return true, nil
	}

	return false, nil
}