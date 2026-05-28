
# Company Registry API — Helm Deployment

This folder contains a Helm chart and Docker image archive for deploying comp-reg-api into Kubernetes without an external Docker registry.

## Package contents

`text
helm-comp-reg-api/
  comp-reg-api_0.1.0.tar
  Chart.yaml
  values.yaml
  README_DEPLOY.md
  templates/
    _helpers.tpl
    deployment.yaml
    service.yaml
    ingress.yaml

Application image

Expected image name:

comp-reg-api:0.1.0

The image archive is included:

comp-reg-api_0.1.0.tar

Application ports

The application listens inside the container on:

8080

The Kubernetes Service exposes it inside the cluster on:

80

Mapping:

Service:80 -> Pod:8080

Endpoints

GET /health
GET /status
GET /swagger
GET /openapi.yaml
GET /api/v1/companies
POST /api/v1/companies/search

Example:

/api/v1/companies?status=Активен
/api/v1/companies?company_name=*Центр*
/api/v1/companies?legal_address=*Москва*
