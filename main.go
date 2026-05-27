package main

import (
	"log"
	"net/http"
)

func main() {
	storage := NewMockCompanyStorage()
	service := NewCompanyService(storage)
	handler := NewHandler(service)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	addr := ":8080"

	log.Printf("company-registry-api started on %s", addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
