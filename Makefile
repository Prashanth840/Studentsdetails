SHELL := /bin/bash

.PHONY: help run fmt vet migrate-new migrate-up migrate-down

help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## run the API server
	go run main.go

fmt: ## run "go fmt" on all packages
	@go fmt ./...

vet: ## run "go vet" on all packages
	@go vet ./...

migrate-new: ## create a new database migration (prompts for a name if NAME is not set)
	@name="$(NAME)"; \
	if [ -z "$$name" ]; then \
		read -p "Enter the name of the new migration: " name; \
	fi; \
	docker run --rm -v "$$(pwd)/migrations:/migrations" migrate/migrate create -ext sql -dir /migrations -seq "$${name// /_}"

migrate-up: ## apply pending migrations (via docker compose)
	docker compose up migrate

migrate-down: ## roll back the last migration
	docker run --rm -v "$$(pwd)/migrations:/migrations" --network studentsdetails_default migrate/migrate \
		-path /migrations -database "postgres://$${DB_USER:-postgres}:$${DB_PASSWORD:-password}@postgres:5432/$${DB_NAME:-prashanth}?sslmode=disable" down 1
