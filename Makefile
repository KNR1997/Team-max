BINARY_NAME=todo
GO=go

.PHONY: build run test test-integration lint fmt vet up down clean help check

## build: Build the application
build:
	$(GO) build ./...

## run: Build and run the application
run: build
	$(GO) run ./cmd/server

## test: Run unit tests with race detection
test:
	$(GO) test -race -count=1 ./...

## test-integration: Run integration tests (needs TEST_DATABASE_URL)
test-integration:
	$(GO) test -race -tags=integration -count=1 ./...

## lint: Run golangci-lint
lint:
	golangci-lint run ./...

## fmt: Format code
fmt:
	gofmt -s -w .
	$(GO) run golang.org/x/tools/cmd/goimports@latest -local todo -w .

## vet: Run go vet
vet:
	$(GO) vet ./...

## check: Run fmt, vet, lint, and tests
check: fmt vet lint test

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

## migrate-up: Apply all migrations
migrate-up:
	migrate -path ./migrations -database "$(DATABASE_URL)" up

## migrate-down: Roll back all migrations
migrate-down:
	migrate -path ./migrations -database "$(DATABASE_URL)" down

## migrate-drop: Drop everything
migrate-drop:
	migrate -path ./migrations -database "$(DATABASE_URL)" drop -f

## migrate-create: Create a new migration (usage: make migrate-create NAME=add_users)
migrate-create:
	@test -n "$(NAME)" || (echo "NAME required" && exit 1)
	migrate create -ext sql -dir ./migrations -seq $(NAME)
