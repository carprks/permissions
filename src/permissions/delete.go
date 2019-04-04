package permissions

import "net/http"

func (pr PermissionRequest) DeleteUser() (p Permission, err error) {
	return p, err
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func (pr PermissionRequest) DeletePermission() (p Permission, err error) {
	return p, err
}

func DeletePermission(w http.ResponseWriter, r *http.Request) {

}
