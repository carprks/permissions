package permissions

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// )
//
// // Create company permission http
// func Create(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
//
// 	j := PermissionRequestHTTP{}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(PermissionResponse{
// 			Error: err,
// 		})
// 		return
// 	}
// 	e := json.Unmarshal(body, &j)
// 	if e != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(PermissionResponse{
// 			Error: e,
// 		})
// 		return
// 	}
//
// 	pr, err := j.ConvertToPermissionRequest()
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(PermissionResponse{
// 			Error: err,
// 		})
// 		return
// 	}
//
// 	resp, err := pr.Create()
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(PermissionResponse{
// 			Error: err,
// 		})
// 		return
// 	}
//
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(PermissionResponse{
// 		Permission: resp,
// 	})
// 	return
// }
//
// // Create company permission
// func (pr PermissionRequest) Create() (p Permission, err error) {
// 	pc := Permission{
// 		ID: pr.getCompanyUUID(),
// 		Name: pr.Name,
// 		AllowedTo: pr.Permission,
// 		Identity: pr.Identity,
// 		Company: true,
// 	}
//
// 	return pc.create()
// }
//
// // CreateUser permission http
// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
//
// 	j := PermissionRequestHTTP{}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(PermissionResponse{
// 			Error: err,
// 		})
// 		return
// 	}
// 	e := json.Unmarshal(body, &j)
// 	if e != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(PermissionResponse{
// 			Error: e,
// 		})
// 		return
// 	}
//
// 	pr, err := j.ConvertToPermissionRequest()
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
//
// 	resp, err := pr.CreateUser()
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(PermissionResponse{
// 			Error: err,
// 		})
// 		return
// 	}
//
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(PermissionResponse{
// 		Permission: resp,
// 	})
// 	return
// }
//
// // CreateUser permission
// func (pr PermissionRequest) CreateUser() (p Permission, err error) {
// 	pc := Permission{
// 		ID: pr.getUserUUID(),
// 		Name: pr.Name,
// 		AllowedTo: pr.Permission,
// 		Identity: pr.Identity,
// 		Company: false,
// 	}
//
// 	return pc.create()
// }
//
// func (pc Permission) create() (Permission, error) {
// 	p := pc
// 	p.Status = PermissionBad
//
// 	store := os.Getenv("DATABASE_DYNAMO")
// 	if store == "true" {
// 		c, err := pc.checkExists()
// 		if err != nil {
// 			return Permission{}, err
// 		}
// 		if !c {
// 			p.Status = PermissionGood
// 			ps, err := p.storeDynamo()
// 			if err != nil {
// 				fmt.Println(fmt.Sprintf("Store Error: %v", err))
// 				p.Status = PermissionBad
//
// 				return p, err
// 			}
// 			return ps, nil
// 		}
// 		return p, nil
// 	}
//
// 	return p, nil
// }
