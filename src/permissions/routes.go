package permissions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateRoute general permission
func CreateRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pr := PermissionRequest{}
	err = json.Unmarshal([]byte(body), &pr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// p, err := pr.Create()
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// json.NewEncoder(w).Encode(p)
}

// RetrieveAllRoute general permission
func RetrieveAllRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(700)
}
// RetrieveRoute general permission
func RetrieveRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(701)
}

// UpdateRoute general permission
func UpdateRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(702)
}

// DeleteRoute general permission
func DeleteRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(703)
}

// CreateUserRoute user permission
func CreateUserRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(800)
}

// RetrieveUserAllRoute user permission
func RetrieveUserAllRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(801)
}

// RetrieveUserRoute user permission
func RetrieveUserRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(802)
}

// UpdateUserRoute user permission
func UpdateUserRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(803)
}

// DeleteUserRoute user permission
func DeleteUserRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(804)
}
