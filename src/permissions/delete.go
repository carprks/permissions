package permissions

import "net/http"

// DeleteUser permission
func (pr PermissionRequest) DeleteUser() (Permission, error) {
	p := Permission{
		ID: pr.getUserUUID(),
	}
	resp, err := p.deleteDynamo()
	if err != nil {
		return Permission{}, err
	}

	return resp, nil
}

// DeleteUser permission http
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

// DeletePermission general
func (pr PermissionRequest) DeletePermission() (Permission, error) {
	p := Permission{
		ID: pr.getCompanyUUID(),
	}
	resp, err := p.deleteDynamo()
	if err != nil {
		return Permission{}, err
	}

	return resp, nil
}

// DeletePermission general http
func DeletePermission(w http.ResponseWriter, r *http.Request) {

}
