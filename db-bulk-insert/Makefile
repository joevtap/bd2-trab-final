.PHONY: setup

all: setup test build run

setup:
	@echo "Starting Docker containers..."
	@docker-compose up -d db
	@echo "Waiting for database to start..."
	@sleep 5
	@echo "Running migrations..."
	@docker-compose up -d migrate
	@echo "Generating GORM models..."
	go run cmd/gen/main.go
	@echo "Done!"

teardown:
	@echo "Stopping Docker containers..."
	@docker-compose down
	@echo "Done!"

purge:
	@echo "Purging DB data..."
	@rm -rf ./db/data
	@echo "Done!"

test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Done!"

build:
	@echo "Building bulk-insert..."
	@go build -o cmd/bulk-insert/bin/bulk-insert cmd/bulk-insert/main.go
	@echo "Done!"

run:
	@./cmd/bulk-insert/bin/bulk-insert