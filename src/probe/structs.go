package probe

// Error model
type Error struct {
	Code    int
	Message string
}

// Healthy model
type Healthy struct {
	Status string `json:"status"`
}
