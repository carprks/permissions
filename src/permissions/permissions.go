package permissions

import (
	"github.com/satori/go.uuid"
)

// Create permission
func (pr PermissionRequest) Create() (p Permission, err error) {
	u := uuid.NewV4()

	p.ID = u.String()
	p.Name = pr.Name
	p.AllowedTo = pr.Permission
	p.User = pr.User

	return p, err
}

// Response for permission
func (p Permission) Response() (pr PermissionResponse) {
	pr.Status = PermissionBad

	if p.Name != "" {
		pr.Status = PermissionGood
	}

	if p.AllowedTo != "" {
		pr.Status = PermissionGood
	}

	return pr
}
