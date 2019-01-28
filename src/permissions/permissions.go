package permissions

import (
	"github.com/satori/go.uuid"
)

func (pr PermissionRequest) Create() (p Permission) {
	u := uuid.NewV4()

	p.Id = u.String()
	p.Name = pr.Name
	p.AllowedTo = pr.Permission
	p.User = pr.User

	return p
}

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
