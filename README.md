# Company Registry API

## Brief Description

Company Registry API is a lightweight demo REST service written in Go.
It provides company search functionality using wildcard filters and returns synthetic Russian company data loaded from a CSV file into memory on startup.

The service includes:
- REST API
- wildcard filtering
- Swagger UI
- health and status endpoints
- Docker Compose deployment
- Kubernetes no-registry deployment package

---

# Краткое описание

Company Registry API — это лёгкий демонстрационный REST-сервис на Go.
Сервис предоставляет поиск компаний с wildcard-фильтрацией и возвращает синтетические данные российских компаний, загружаемые из CSV-файла в память при старте.

Сервис включает:
- REST API
- wildcard-фильтрацию
- Swagger UI
- health/status endpoints
- Docker Compose deployment
- Kubernetes deployment

---

# Project Structure

```text
comp_reg_api/
  *.go
  go.mod
  openapi.yaml
  swagger-ui/
  data/

  docker_compose/

  k8s-comp-reg-api/

  helm-comp-reg-api/
  
```

---

# How To Start The Service

## Run locally

```powershell
go run .
```

Service URLs:

```text
http://localhost:8080/swagger
http://localhost:8080/status
http://localhost:8080/health
```

---

## Run with Docker Compose

```powershell
docker compose -f comp-reg-api.yaml up --build
```

Stop:

```powershell
docker compose -f comp-reg-api.yaml down
```

---

# Kubernetes Deployment

The `k8s-comp-reg-api` folder contains a complete no-registry deployment package - not tested.

The `helm-comp-reg-api` has helm chart, uses image registry - tested.

---

# API Endpoints

## Swagger UI

```text
GET /swagger
```

Example:

```text
http://localhost:8080/swagger
```

---

## OpenAPI YAML

```text
GET /openapi.yaml
```

Example:

```text
http://localhost:8080/openapi.yaml
```

---

## Health Check

```text
GET /health
```

Example:

```text
http://localhost:8080/health
```

Response:

```json
{
  "status": "ok"
}
```

---

## Service Status

```text
GET /status
```

Example:

```text
http://localhost:8080/status
```

Example response:

```json
{
  "service": "company-registry-api",
  "status": "running",
  "records_count": 150
}
```

---

## Search Companies (GET)

```text
GET /api/v1/companies
```

Supported filters:
- company_name
- legal_form
- legal_address
- status

Wildcard support:
- `*` — any number of symbols
- `?` — single symbol

Examples:

### All companies

```text
http://localhost:8080/api/v1/companies
```

### By status

```text
http://localhost:8080/api/v1/companies?status=Активен
```

### By city

```text
http://localhost:8080/api/v1/companies?legal_address=*Москва*
```

### By company name

```text
http://localhost:8080/api/v1/companies?company_name=*Центр*
```

### Combined filters

```text
http://localhost:8080/api/v1/companies?legal_form=ООО&status=Активен
```

---

## Search Companies (POST)

```text
POST /api/v1/companies/search
```

Content-Type:

```text
application/json
```

Example body:

```json
{
  "company_name": "*Логистика*",
  "status": "Активен"
}
```

Example curl:

```bash
curl -X POST http://localhost:8080/api/v1/companies/search \
  -H "Content-Type: application/json" \
  -d '{
    "company_name": "*Логистика*",
    "status": "Активен"
  }'
```

---

# Architecture

The service uses a simple layered architecture.

```text
HTTP Handlers
      ↓
Business Service
      ↓
Storage Layer
      ↓
CSV File
```

---

# File/Module Description

## main.go

Application entry point.
Creates storage, service, handlers and starts HTTP server.

---

## models.go

Contains main data structures:
- Company
- CompanyFilter

---

## storage.go

Responsible for:
- loading CSV data
- storing companies in memory
- storage abstraction

---

## service.go

Contains business logic:
- filtering
- wildcard matching
- search orchestration

---

## wildcard.go

Converts wildcard patterns into Go regular expressions.

Examples:

```text
*Москва*
ООО*
*Логистика
```

---

## handlers.go

HTTP handlers:
- REST API
- health endpoint
- status endpoint
- JSON serialization

---

## swagger.go

Swagger/OpenAPI routes and static Swagger UI hosting.

---

## openapi.yaml

OpenAPI specification used by Swagger UI.

---

## swagger-ui/

Contains local Swagger UI static files.

---

## data/companies.csv

Synthetic company dataset loaded into memory on service startup.

---

# Technologies

- Go
- net/http
- Docker
- Docker Compose
- Kubernetes
- Swagger/OpenAPI

---

# Notes

This project is intentionally simple and lightweight for:
- demos
- integration testing
- API mocking
- presales
