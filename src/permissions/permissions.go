package permissions

import (
	// "errors"
	"fmt"
	"github.com/satori/go.uuid"
	// "strings"
)

// PermissionResponse struct
type PermissionResponse struct {
	Permissions Permissions `json:"permissions,omitempty"`
	Error       error       `json:"error,omitempty"`
}

// Permissions struct
type Permissions struct {
	ID          string       `json:"id"`
	Identity    string       `json:"identity"`
	Permissions []Permission `json:"permissions"`
	Company bool `json:"company"`
}

// Permission struct
type Permission struct {
	Name   string `json:"name"`
	Action string `json:"action"`
}

// PermissionRequestHTTP struct
// type PermissionRequestHTTP struct {
// 	Permission string `json:"permission"`
// 	Identity   string `json:"identity"`
// }

// PermissionRequest struct
type PermissionRequest struct {
	Identity    string
	Permissions []Permission
}

// PermissionRequestUpdateHTTP struct
// type PermissionRequestUpdateHTTP struct {
// 	Old PermissionRequestHTTP `json:"old"`
// 	New PermissionRequestHTTP `json:"new"`
// }

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

// checkExists
// func (p Permission) checkExists() (bool, error) {
// 	resp, err := p.retrieveDynamo()
// 	if err != nil {
// 		if err.Error() == "no permission entry" {
// 			return false, nil
// 		}
//
// 		return false, err
// 	}
//
// 	// it has an id so it must exist
// 	if resp.ID != "" {
// 		return true, nil
// 	}
//
// 	return false, nil
// }

// func (prh PermissionRequestHTTP) ConvertToPermission(company bool) (Permissions, error) {
// 	pc := Permissions{}
// 	pr, err := prh.ConvertToPermissionRequest()
// 	if err != nil {
// 		return pc, err
// 	}
//
// 	if company {
// 		pc = Permissions{
// 			ID:          pr.getCompanyUUID(),
// 			Permissions: pr.Permissions,
// 			Identity:    pr.Identity,
// 			Company:     company,
// 		}
// 	} else {
// 		pc = Permissions{
// 			ID:          pr.getUserUUID(),
// 			Permissions: pr.Permissions,
// 			Identity:    pr.Identity,
// 			Company:     company,
// 		}
// 	}
//
// 	return pc, nil
// }

// func (prh PermissionRequestHTTP) ConvertToPermissionRequest() (PermissionRequest, error) {
// 	pr := PermissionRequest{}
//
// 	if prh.Permission == "*" {
// 		pr = PermissionRequest{
// 			Action:   "*",
// 			Name:     "*",
// 			Identity: prh.Identity,
// 		}
// 	} else {
// 		split := strings.Split(prh.Permission, ":")
// 		if len(split) != 2 {
// 			return pr, errors.New("invalid permissions")
// 		}
// 		pr = PermissionRequest{
// 			Action:   split[1],
// 			Name:     split[0],
// 			Identity: prh.Identity,
// 		}
// 	}
//
// 	return pr, nil
// }

func (p Permissions) MapPermissions() map[string]string {
	ret := map[string]string{}

	if len(p.Permissions) >= 1 {
		for _, perm := range p.Permissions {
			ret[perm.Name] = perm.Action
		}
	}

	return ret
}
