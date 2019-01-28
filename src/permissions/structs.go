package permissions

type Permission struct {
	Id        string
	Name      string
	AllowedTo string
	User      string
}

type PermissionRequest struct {
	Name       string `json:"name"`
	Permission string `json:"permission"`
	User       string `json:"user"`
}

type PermissionResponse struct {
	Status PermissionStatus `json:"status"`
}

type PermissionStatus string

const (
	PermissionBad  PermissionStatus = "Bad"
	PermissionGood PermissionStatus = "Good"
)
