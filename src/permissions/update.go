package permissions

import "net/http"

func (pr PermissionRequest) UpdateUser() (p Permission, err error) {
	return p, err
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (pr PermissionRequest) UpdatePermission() (p Permission, err error) {
	return p, err
}

func UpdatePermission(w http.ResponseWriter, r *http.Request) {

}
