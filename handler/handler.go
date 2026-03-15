package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/your-username/health-check-monitoring-service/model"
	"github.com/your-username/health-check-monitoring-service/service"
)

// Handler represents the application handler
type Handler interface {
	GetHealthChecks(w http.ResponseWriter, r *http.Request)
	GetHealthCheck(w http.ResponseWriter, r *http.Request)
	CreateHealthCheck(w http.ResponseWriter, r *http.Request)
	UpdateHealthCheck(w http.ResponseWriter, r *http.Request)
	DeleteHealthCheck(w http.ResponseWriter, r *http.Request)
}

// NewHandler returns a new application handler
func NewHandler(srv service.Service) Handler {
	return &handler{srv: srv}
}

type handler struct {
	srv service.Service
}

func (h *handler) GetHealthChecks(w http.ResponseWriter, r *http.Request) {
	healthChecks, err := h.srv.GetHealthChecks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(healthChecks)
}

func (h *handler) GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	healthCheck, err := h.srv.GetHealthCheck(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(healthCheck)
}

func (h *handler) CreateHealthCheck(w http.ResponseWriter, r *http.Request) {
	var healthCheck model.HealthCheck
	if err := json.NewDecoder(r.Body).Decode(&healthCheck); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.srv.CreateHealthCheck(&healthCheck); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) UpdateHealthCheck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var healthCheck model.HealthCheck
	if err := json.NewDecoder(r.Body).Decode(&healthCheck); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.srv.UpdateHealthCheck(&healthCheck); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) DeleteHealthCheck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.srv.DeleteHealthCheck(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func atoi(s string) (int, error) {
	var i int
	if _, err := fmt.Sscan(s, &i); err != nil {
		return 0, err
	}
	return i, nil
}