FROM golang:1.22-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY *.go ./

RUN go build -ldflags="-s -w" -o comp-reg-api .

FROM alpine:3.20

WORKDIR /app

COPY --from=build /src/comp-reg-api .
COPY data ./data
COPY openapi.yaml .
COPY swagger-ui ./swagger-ui

ENV PORT=8080
ENV COMPANIES_FILE=data/companies.csv

EXPOSE 8080

CMD ["./comp-reg-api"]