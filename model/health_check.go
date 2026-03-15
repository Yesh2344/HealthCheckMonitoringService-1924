package model

// HealthCheck represents a health check
// Updated - v2.0.7
// Updated - v8.2.8
type HealthCheck struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	URL  string `json:"url" db:"url"`
}