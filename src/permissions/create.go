package permissions

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"os"
)

// Create permission
func (pr PermissionRequest) Create() (p Permission, err error) {
	u := uuid.NewV5(uuid.NamespaceURL, fmt.Sprintf("https://permissions.carprk.com/company/%s", pr.User))

	return pr.create(u.String())
}

// Create HTTP
func Create(w http.ResponseWriter, r *http.Request) {

}

// CreateUser HTTP
func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (pr PermissionRequest) CreateUser() (p Permission, err error) {
	u := uuid.NewV5(uuid.NamespaceDNS, fmt.Sprintf("https://permissions.carprk.com/user/%s", pr.User))

	return pr.create(u.String())
}

func (pr PermissionRequest) create(uuid string) (Permission, error) {
	p := Permission{
		ID: uuid,
		Name: pr.Name,
		AllowedTo: pr.Permission,
		User: pr.User,
		Status: PermissionBad,
	}

	if os.Getenv("DATABASE_DYNAMO") == "true" {
		p, err := p.storeDynamo()
		if err != nil {
			fmt.Println(fmt.Sprintf("Store User Error: %v", err))

			return p, err
		}
	}

	return p, nil
}