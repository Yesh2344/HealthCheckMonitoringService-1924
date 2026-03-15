package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/your-username/health-check-monitoring-service/model"
)

// Repository represents the database repository
type Repository interface {
	GetHealthChecks() ([]model.HealthCheck, error)
	GetHealthCheck(id int) (*model.HealthCheck, error)
	CreateHealthCheck(healthCheck *model.HealthCheck) error
	UpdateHealthCheck(healthCheck *model.HealthCheck) error
	DeleteHealthCheck(id int) error
}

// NewRepository returns a new database repository
func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) GetHealthChecks() ([]model.HealthCheck, error) {
	var healthChecks []model.HealthCheck
	if err := r.db.Select(&healthChecks, "SELECT * FROM health_checks"); err != nil {
		return nil, err
	}
	return healthChecks, nil
}

func (r *repository) GetHealthCheck(id int) (*model.HealthCheck, error) {
	var healthCheck model.HealthCheck
	if err := r.db.Get(&healthCheck, "SELECT * FROM health_checks WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &healthCheck, nil
}

func (r *repository) CreateHealthCheck(healthCheck *model.HealthCheck) error {
	if _, err := r.db.NamedExec("INSERT INTO health_checks (name, url) VALUES (:name, :url)", healthCheck); err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateHealthCheck(healthCheck *model.HealthCheck) error {
	if _, err := r.db.NamedExec("UPDATE health_checks SET name = :name, url = :url WHERE id = :id", healthCheck); err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteHealthCheck(id int) error {
	if _, err := r.db.Exec("DELETE FROM health_checks WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}