package permissions

import "net/http"

// UpdateUser update the user permission
func (pr PermissionRequestUpdate) UpdateUser() (Permission, error) {
	old := Permission{
		ID: pr.Old.getCompanyUUID(),
		Name: pr.Old.Name,
		AllowedTo: pr.Old.Permission,
		User: pr.Old.User,
		Company: false,
	}
	found, err := old.retrieve()
	if err != nil {
		return Permission{}, err
	}
	if found.Status == PermissionGood {
		old.Status = PermissionBad
		_, err := old.update()
		if err != nil {
			return Permission{}, err
		}

		newP := Permission{
			ID: pr.New.getUserUUID(),
			Name: pr.New.Name,
			AllowedTo: pr.New.Permission,
			User: pr.New.User,
			Company: false,
		}
		resp, err := newP.create()
		if err != nil {
			return Permission{}, err
		}

		return resp, nil
	}

	return Permission{}, nil
}

// UpdateUser update the user permission http
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// UpdatePermission update the permission
func (pr PermissionRequestUpdate) UpdatePermission() (Permission, error) {
	old := Permission{
		ID: pr.Old.getCompanyUUID(),
		Name: pr.Old.Name,
		AllowedTo: pr.Old.Permission,
		User: pr.Old.User,
		Company: true,
	}
	found, err := old.retrieve()
	if err != nil {
		return Permission{}, err
	}
	if found.Status == PermissionGood {
		old.Status = PermissionBad
		_, err := old.update()
		if err != nil {
			return Permission{}, err
		}

		newP := Permission{
			ID: pr.New.getCompanyUUID(),
			Name: pr.New.Name,
			AllowedTo: pr.New.Permission,
			User: pr.New.User,
			Company: true,
		}
		resp, err := newP.create()
		if err != nil {
			return Permission{}, err
		}

		return resp, nil
	}

	return Permission{}, nil
}

// UpdatePermission update the permission http
func UpdatePermission(w http.ResponseWriter, r *http.Request) {

}

func (p Permission) update() (Permission, error) {
	resp, err := p.updateDynamo()
	if err != nil {
		return Permission{}, err
	}

	return resp, nil
}