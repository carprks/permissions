package permissions

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

// Delete the permission
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	identity := chi.URLParam(r, "identityID")
	p := Permission{
		Identity: identity,
	}

	resp, err := p.DeleteEntry()
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(PermissionResponse{
		Permission: resp,
	})
	if err != nil {
		ErrorResponse(w, err)
	}
}