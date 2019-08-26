package permissions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Create the permission
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := Permissions{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	if len(p.Permissions) == 0 {
		err = fmt.Errorf("need at least 1 permission")
		ErrorResponse(w, err)
		return
	}

	resp, err := p.CreateEntry()
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(PermissionResponse{
		Permissions: resp,
	})
	if err != nil {
		ErrorResponse(w, err)
	}
}