package model

// HealthCheck represents a health check
type HealthCheck struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	URL  string `json:"url" db:"url"`
}