BINARY_NAME=todo
GO=go

.PHONY: build run test test-integration lint fmt up down clean help

## build: Build the application
build:
	$(GO) build -o bin/$(BINARY_NAME) ./cmd/server

## run: Build and run the application
run: build
	./bin/$(BINARY_NAME)

## test: Run unit tests with race detection
test:
	$(GO) test -race -coverprofile=coverage.out ./...

## test-integration: Run integration tests (needs TEST_DATABASE_URL)
test-integration:
	$(GO) test -race -tags=integration -count=1 ./internal/task/

## lint: Run golangci-lint
lint:
	golangci-lint run ./...

## fmt: Format and vet
fmt:
	gofmt -s -w cmd internal
	$(GO) vet ./...

## up: Start Postgres
up:
	docker compose up -d --wait

## down: Stop Postgres
down:
	docker compose down

## clean: Clean build artifacts
clean:
	rm -rf bin/ coverage.out

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'
