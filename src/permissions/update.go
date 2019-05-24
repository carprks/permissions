package permissions

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	identity := chi.URLParam(r, "identityID")
	p := Permission{
		Identity: identity,
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	req := Permission{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	up, err := p.UpdateEntry(req)
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(PermissionResponse{
		Permission: up,
	})
	if err != nil {
		ErrorResponse(w, err)
		return
	}
}