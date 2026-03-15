package service

import (
	"github.com/your-username/health-check-monitoring-service/model"
	"github.com/your-username/health-check-monitoring-service/repository"
)

// Service represents the application service
type Service interface {
	GetHealthChecks() ([]model.HealthCheck, error)
	GetHealthCheck(id int) (*model.HealthCheck, error)
	CreateHealthCheck(healthCheck *model.HealthCheck) error
	UpdateHealthCheck(healthCheck *model.HealthCheck) error
	DeleteHealthCheck(id int) error
}

// NewService returns a new application service
func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

type service struct {
	repo repository.Repository
}

func (s *service) GetHealthChecks() ([]model.HealthCheck, error) {
	return s.repo.GetHealthChecks()
}

func (s *service) GetHealthCheck(id int) (*model.HealthCheck, error) {
	return s.repo.GetHealthCheck(id)
}

func (s *service) CreateHealthCheck(healthCheck *model.HealthCheck) error {
	return s.repo.CreateHealthCheck(healthCheck)
}

func (s *service) UpdateHealthCheck(healthCheck *model.HealthCheck) error {
	return s.repo.UpdateHealthCheck(healthCheck)
}

func (s *service) DeleteHealthCheck(id int) error {
	return s.repo.DeleteHealthCheck(id)
}