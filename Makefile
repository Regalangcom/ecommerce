help:
	@echo "Available commands:"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make dev          - Run the application in development mode"
	@echo "  make lint         - Run linters"
	@echo "  make format       - Run formatting"
	@echo "  make migrate-up   - Run all pending migrations"
	@echo "  make migrate-down - Rollback the last migration"
	@echo "  make help         - Display this help message"

build:
	go build -o bin/app ./cmd/api

run: 
	go run ./cmd/api

dev:
	go run ./cmd/api

lint: format
	golangci-lint run ./...

format:
	@gofmt -s -w .
	@goimports -w .


migrate-up:
	migrate -path db/migrations -database "postgresql://ucommerce:admin1234@localhost:5432/ecommerce-shop?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgresql://ucommerce:admin1234@localhost:5432/ecommerce-shop?sslmode=disable" down

docker-up:
	docker compose -f docker/docker-compose.yml up -d 

docker-down:
	docker compose -f docker/docker-compose.yml down 