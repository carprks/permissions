package permissions

import (
	"fmt"
	"net/http"
	"os"
)

// Create company permission http
func Create(w http.ResponseWriter, r *http.Request) {

}

// Create company permission
func (pr PermissionRequest) Create() (p Permission, err error) {
	pc := Permission{
		ID: pr.getCompanyUUID(),
		Name: pr.Name,
		AllowedTo: pr.Permission,
		User: pr.User,
		Company: true,
	}

	return pc.create()
}

// CreateUser permission http
func CreateUser(w http.ResponseWriter, r *http.Request) {

}

// CreateUser permission
func (pr PermissionRequest) CreateUser() (p Permission, err error) {
	pc := Permission{
		ID: pr.getUserUUID(),
		Name: pr.Name,
		AllowedTo: pr.Permission,
		User: pr.User,
		Company: false,
	}

	return pc.create()
}

func (pc Permission) create() (Permission, error) {
	p := pc
	p.Status = PermissionBad

	store := os.Getenv("DATABASE_DYNAMO")
	if store == "true" {
		c, err := pc.checkExists()
		if err != nil {
			return Permission{}, err
		}
		if !c {
			p.Status = PermissionGood
			p, err := p.storeDynamo()
			if err != nil {
				fmt.Println(fmt.Sprintf("Store Error: %v", err))
				p.Status = PermissionBad

				return p, err
			}
		}
		return p, nil
	}

	return p, nil
}