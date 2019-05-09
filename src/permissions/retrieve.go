package permissions

// import (
// 	"net/http"
// 	"os"
// )
//
// // RetrieveAll retrieve all permissions
// func (pr PermissionRequest) RetrieveAll() (p Permission, err error) {
// 	return p, err
// }
//
// // RetrieveAll retrieve all permission http
// func RetrieveAll(w http.ResponseWriter, r *http.Request) {
//
// }
//
// // RetrieveAllUsers retrieve all user permissions
// func (pr PermissionRequest) RetrieveAllUsers() (p Permission, err error) {
// 	return p, err
// }
//
// // RetrieveAllUsers retrieve all user permissions http
// func RetrieveAllUsers(w http.ResponseWriter, r *http.Request) {
//
// }
//
// // RetrievePermissions retrieve the permission
// func (pr PermissionRequest) RetrievePermissions() (p Permission, err error) {
// 	pc := Permission{
// 		ID: pr.getCompanyUUID(),
// 		Name: pr.Name,
// 		AllowedTo: pr.Permission,
// 		Identity: pr.Identity,
// 		Company: true,
// 		Status: PermissionGood,
// 	}
//
// 	return pc.retrieve()
// }
//
// // RetrievePermissions retrieve the permission http
// func RetrievePermissions(w http.ResponseWriter, r *http.Request){
//
// }
//
// // RetrieveUser retrieve the user permission
// func (pr PermissionRequest) RetrieveUser() (p Permission, err error) {
// 	pc := Permission{
// 		ID: pr.getUserUUID(),
// 		Name: pr.Name,
// 		AllowedTo: pr.Permission,
// 		Identity: pr.Identity,
// 		Company: false,
// 	}
//
// 	return pc.retrieve()
// }
//
// // RetrieveUser retrieve the user permission http
// func RetrieveUser(w http.ResponseWriter, r *http.Request) {
// }
//
// func (pc Permission) retrieve() (Permission, error) {
// 	p := pc
// 	p.Status = PermissionBad
//
// 	store := os.Getenv("DATABASE_DYNAMO")
// 	if store == "true" {
// 		pa, err := p.retrieveAstrixDynamo()
// 		if err != nil {
// 			// fmt.Println(fmt.Sprintf("Store Retrieve Astrix Error: %v", err))
// 			return p, err
// 		}
// 		return pa, nil
// 	}
//
// 	return p, nil
// }