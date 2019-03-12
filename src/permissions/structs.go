package permissions

// Permission struct
type Permission struct {
	ID        string
	Name      string
	AllowedTo string
	User      string
}

// PermissionRequest struct
type PermissionRequest struct {
	Name       string `json:"name"`
	Permission string `json:"permission"`
	User       string `json:"user"`
}

// PermissionResponse struct
type PermissionResponse struct {
	Status PermissionStatus `json:"status"`
}

// PermissionStatus string
type PermissionStatus string

const (
	// PermissionBad status of the permission
	PermissionBad  PermissionStatus = "Bad"

	// PermissionGood status of the permission
	PermissionGood PermissionStatus = "Good"
)
