.PHONY: run build migrate worker test fmt vet tidy clean db-schema

# Start the API locally.
run:
	go run ./cmd/server

# Compile the API binary into bin/.
build:
	mkdir -p bin
	go build -o bin/server ./cmd/server

# Run database migrations.
migrate:
	go run ./cmd/migrate

# Start the background worker.
worker:
	go run ./cmd/worker

# Run all package tests.
test:
	go test ./...

# Format all Go source files.
fmt:
	gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

# Run static checks.
vet:
	go vet ./...

# Synchronize Go module dependencies.
tidy:
	go mod tidy

# Load the database schema. Requires MYSQL_USER, MYSQL_PASSWORD, and MYSQL_DATABASE.
db-schema:
	mysql -u "$${MYSQL_USER}" -p"$${MYSQL_PASSWORD}" "$${MYSQL_DATABASE}" < internal/database/schema.sql

# Remove locally built binaries.
clean:
	rm -rf bin
