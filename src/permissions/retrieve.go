package permissions

import (
  "encoding/json"
  "github.com/go-chi/chi"
  "net/http"
  "strings"
)

// Retrieve get all the permissions
func Retrieve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	identity := chi.URLParam(r, "identityID")
	p := Permission{
		Identity: identity,
	}

	perms, err := p.RetrieveEntry()
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(PermissionResponse{
		Permission: perms,
	})
	if err != nil {
		ErrorResponse(w, err)
	}
}

// Allowed can we do this action
func Allowed(w http.ResponseWriter, r *http.Request) {
	identity := chi.URLParam(r, "identityID")
	permission := chi.URLParam(r, "permission")
	permReq := strings.Split(permission, ":")
	allowed := false

	p := Permission{
		Identity: identity,
	}
	perms, err := p.RetrieveEntry()
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	for _, perm := range perms.Permissions {
		if perm.Name == PermissionAstrix && perm.Action == PermissionAstrix {
			allowed = true
		}

		if permReq[0] == perm.Name && permReq[1] == perm.Action {
			allowed = true
		}
	}

	if !allowed {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}
