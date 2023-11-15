.PHONY: all db app setup teardown purge build-win build-darwin build-linux

all: setup app

db:
	@echo "Starting Docker containers..."
	@sudo docker-compose up -d db
	@echo "Done!"

app:
	@echo "Starting Docker containers..."
	@sudo docker-compose up -d app
	@echo "Done!"

setup:
	@echo "Starting Docker containers..."
	@sudo docker-compose up -d db
	@echo "Waiting for database to start..."
	@sleep 5
	@echo "Running migrations..."
	@sudo docker-compose up -d migrate
	@echo "Done!"

teardown:
	@echo "Stopping Docker containers..."
	@sudo docker-compose down
	@echo "Done!"

purge:
	@echo "Purging DB data..."
	@sudo rm -rf ./db/data
	@echo "Done!"

build-win:
	@echo "Building bulk-insert..."
	@docker run -e GOOS=windows -e GOARCH=amd64 --rm -v .:/usr/src/app -w /usr/src/app golang:1.21 go build -o ./cmd/bulk-insert/bin/bulk-insert-windows-amd64.exe ./cmd/bulk-insert/main.go
	@echo "Done!"

build-darwin:
	@echo "Building bulk-insert..."
	@docker run -e GOOS=darwin -e GOARCH=amd64 --rm -v .:/usr/src/app -w /usr/src/app golang:1.21 go build -o ./cmd/bulk-insert/bin/bulk-insert-darwin-amd64 ./cmd/bulk-insert/main.go
	@echo "Done!"

build-linux:
	@echo "Building bulk-insert..."
	@docker run -e GOOS=linux -e GOARCH=amd64 --rm -v .:/usr/src/app -w /usr/src/app golang:1.21 go build -o ./cmd/bulk-insert/bin/bulk-insert-linux-amd64 ./cmd/bulk-insert/main.go
	@echo "Done!"
