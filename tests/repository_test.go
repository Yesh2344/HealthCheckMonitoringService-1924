package tests

import (
	"testing"

	"github.com/your-username/health-check-monitoring-service/database"
	"github.com/your-username/health-check-monitoring-service/model"
	"github.com/jmoiron/sqlx"
)

func TestRepositoryGetHealthChecks(t *testing.T) {
	// Connect to test database
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=your-username password=your-password dbname=your-database sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// Create repository
	repo := database.NewRepository(db)

	// Get health checks
	healthChecks, err := repo.GetHealthChecks()
	if err != nil {
		t.Fatal(err)
	}

	// Check if health checks are not empty
	if len(healthChecks) == 0 {
		t.Errorf("expected health checks to not be empty")
	}
}

func TestRepositoryGetHealthCheck(t *testing.T) {
	// Connect to test database
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=your-username password=your-password dbname=your-database sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// Create repository
	repo := database.NewRepository(db)

	// Get health check
	healthCheck, err := repo.GetHealthCheck(1)
	if err != nil {
		t.Fatal(err)
	}

	// Check if health check is not nil
	if healthCheck == nil {
		t.Errorf("expected health check to not be nil")
	}
}

func TestRepositoryCreateHealthCheck(t *testing.T) {
	// Connect to test database
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=your-username password=your-password dbname=your-database sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// Create repository
	repo := database.NewRepository(db)

	// Create health check
	healthCheck := &model.HealthCheck{
		Name: "test",
		URL:  "http://example.com",
	}

	// Create health check
	if err := repo.CreateHealthCheck(healthCheck); err != nil {
		t.Fatal(err)
	}

	// Get health check
	createdHealthCheck, err := repo.GetHealthCheck(1)
	if err != nil {
		t.Fatal(err)
	}

	// Check if health check matches
	if createdHealthCheck.Name != healthCheck.Name || createdHealthCheck.URL != healthCheck.URL {
		t.Errorf("expected health check to match")
	}
}

func TestRepositoryUpdateHealthCheck(t *testing.T) {
	// Connect to test database
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=your-username password=your-password dbname=your-database sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// Create repository
	repo := database.NewRepository(db)

	// Get health check
	healthCheck, err := repo.GetHealthCheck(1)
	if err != nil {
		t.Fatal(err)
	}

	// Update health check
	healthCheck.Name = "updated"
	healthCheck.URL = "http://example.com/updated"

	// Update health check
	if err := repo.UpdateHealthCheck(healthCheck); err != nil {
		t.Fatal(err)
	}

	// Get updated health check
	updatedHealthCheck, err := repo.GetHealthCheck(1)
	if err != nil {
		t.Fatal(err)
	}

	// Check if health check matches
	if updatedHealthCheck.Name != healthCheck.Name || updatedHealthCheck.URL != healthCheck.URL {
		t.Errorf("expected health check to match")
	}
}

func TestRepositoryDeleteHealthCheck(t *testing.T) {
	// Connect to test database
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=your-username password=your-password dbname=your-database sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// Create repository
	repo := database.NewRepository(db)

	// Get health check
	healthCheck, err := repo.GetHealthCheck(1)
	if err != nil {
		t.Fatal(err)
	}

	// Delete health check
	if err := repo.DeleteHealthCheck(1); err != nil {
		t.Fatal(err)
	}

	// Get health check
	_, err = repo.GetHealthCheck(1)
	if err == nil {
		t.Errorf("expected health check to be deleted")
	}
}