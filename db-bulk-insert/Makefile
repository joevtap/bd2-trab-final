.PHONY: setup teardown purge test build run

all: setup test build run

setup:
	@echo "Starting Docker containers..."
	@sudo docker-compose up -d db
	@echo "Waiting for database to start..."
	@sleep 5
	@echo "Running migrations..."
	@sudo docker-compose up -d migrate
	@echo "Generating GORM models..."
	go run cmd/gen/main.go
	@echo "Done!"

teardown:
	@echo "Stopping Docker containers..."
	@sudo docker-compose down
	@echo "Done!"

purge:
	@echo "Purging DB data..."
	@sudo rm -rf ./db/data
	@echo "Done!"

test:
	@echo "Running tests..."
	@go test -v ./internal/...
	@go test -v ./gorm/...
	@go test -v ./cmd/...
	@echo "Done!"

build:
	@echo "Building bulk-insert..."
	@go build -o cmd/bulk-insert/bin/bulk-insert cmd/bulk-insert/main.go
	@echo "Done!"

run:
	@./cmd/bulk-insert/bin/bulk-insert