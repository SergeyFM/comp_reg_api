package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := getEnv("PORT", "8080")
	dataFile := getEnv("COMPANIES_FILE", "data/companies.csv")

	storage, err := NewFileCompanyStorage(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	service := NewCompanyService(storage)
	handler := NewHandler(service)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)
	RegisterSwaggerRoutes(mux)

	addr := ":" + port

	log.Printf("company-registry-api started on %s", addr)
	log.Printf("loaded companies: %d", storage.Count())

	err = http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
