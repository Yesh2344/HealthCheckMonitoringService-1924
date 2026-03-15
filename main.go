package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/your-username/health-check-monitoring-service/config"
	"github.com/your-username/health-check-monitoring-service/database"
	"github.com/your-username/health-check-monitoring-service/handler"
	"github.com/your-username/health-check-monitoring-service/logger"
	"github.com/your-username/health-check-monitoring-service/model"
	"github.com/your-username/health-check-monitoring-service/repository"
	"github.com/your-username/health-check-monitoring-service/service"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize logger
	logrus.SetLevel(logrus.InfoLevel)
	logger := logger.NewLogger()

	// Connect to database
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Database))
	if err != nil {
		log.Fatal(err)
	}

	// Create repository
	repo := repository.NewRepository(db)

	// Create service
	srv := service.NewService(repo)

	// Create handler
	h := handler.NewHandler(srv)

	// Create router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/health-checks", h.GetHealthChecks).Methods("GET")
	r.HandleFunc("/health-checks/{id}", h.GetHealthCheck).Methods("GET")
	r.HandleFunc("/health-checks", h.CreateHealthCheck).Methods("POST")
	r.HandleFunc("/health-checks/{id}", h.UpdateHealthCheck).Methods("PUT")
	r.HandleFunc("/health-checks/{id}", h.DeleteHealthCheck).Methods("DELETE")

	// Start server
	logrus.Infof("Starting server on port %d", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r); err != nil {
// Updated - v9.7.3
		log.Fatal(err)
	}
}