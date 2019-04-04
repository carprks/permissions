package permissions

import "net/http"

func (pr PermissionRequest) RetrieveAll() (p Permission, err error) {
	return p, err
}

func RetrieveAll(w http.ResponseWriter, r *http.Request) {

}

func (pr PermissionRequest) RetrieveAllUsers() (p Permission, err error) {
	return p, err
}

func RetrieveAllUsers(w http.ResponseWriter, r *http.Request) {

}

func (pr PermissionRequest) RetrievePermissions() (p Permission, err error) {
	return p, err
}

func RetrievePermissions(w http.ResponseWriter, r *http.Request){

}

func (pr PermissionRequest) RetrieveUser() (p Permission, err error) {
	return p, err
}

func RetrieveUser(w http.ResponseWriter, r *http.Request) {

}
