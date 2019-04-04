package permissions

// Permission struct
type Permission struct {
	ID        string
	Name      string
	AllowedTo string
	User      string
	Status    PermissionStatus
}

// PermissionRequest struct
type PermissionRequest struct {
	Name       string `json:"name"`
	Permission string `json:"permission"`
	User       string `json:"user"`
}

// PermissionStatus string
type PermissionStatus string

const (
	// PermissionBad status of the permission
	PermissionBad  PermissionStatus = "Bad"

	// PermissionGood status of the permission
	PermissionGood PermissionStatus = "Good"
)
