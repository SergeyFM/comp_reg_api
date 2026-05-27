package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {
	service   *CompanyService
	startedAt time.Time
}

func NewHandler(service *CompanyService) *Handler {
	return &Handler{
		service:   service,
		startedAt: time.Now(),
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /status", h.Status)
	mux.HandleFunc("GET /api/v1/companies", h.GetCompanies)
	mux.HandleFunc("POST /api/v1/companies/search", h.SearchCompanies)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	writeJSON(w, http.StatusOK, map[string]any{
		"service":        "company-registry-api",
		"status":         "running",
		"started_at":     h.startedAt.Format(time.RFC3339),
		"uptime_sec":     int(now.Sub(h.startedAt).Seconds()),
		"records_count":  h.service.Count(),
		"local_datetime": now.Format("2006-01-02 15:04:05"),
	})
}

func (h *Handler) GetCompanies(w http.ResponseWriter, r *http.Request) {
	filter := CompanyFilter{
		CompanyName: r.URL.Query().Get("company_name"),
	}

	result, err := h.service.FindCompanies(filter)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func (h *Handler) SearchCompanies(w http.ResponseWriter, r *http.Request) {
	var filter CompanyFilter

	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := h.service.FindCompanies(filter)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func writeJSON(w http.ResponseWriter, statusCode int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeJSON(w, statusCode, map[string]string{
		"error": message,
	})
}
