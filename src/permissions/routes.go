package permissions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	type Healthy struct {
		Status string `json:"status"`
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/health+json")
	j, _ := json.Marshal(Healthy{
		Status: "pass",
	})
	w.Write(j)
	return
}

// General
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

	p := pr.Create()
	json.NewEncoder(w).Encode(p.Response())
}

func RetrieveAllRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(700)
}
func RetrieveRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(701)
}

func UpdateRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(702)
}

func DeleteRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(703)
}

// User
func CreateUserRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(800)
}

func RetrieveUserAllRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(801)
}

func RetrieveUserRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(802)
}

func UpdateUserRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(803)
}

func DeleteUserRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(804)
}
