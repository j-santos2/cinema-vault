include .envrc

# ======================================================= #
# AUX
# ======================================================= #
## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ======================================================= #
# DEV
# ======================================================= #
## run/api: run cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api -db-dsn=${CINEMAVAULT_DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo "Creating migration files for ${name}..."
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo "Running migrations..." 
	migrate -path ./migrations -database ${CINEMAVAULT_DB_DSN} up

# ======================================================= #
# QUALITY
# ======================================================= #
## audit: tidy and verify dependencies, format and check code, finally run tests
.PHONY: audit
audit:
	@echo "Tidying and verifying module dependencies..."
	go mod tidy
	go mod verify
	@echo "Formatting code..."
	go fmt ./...
	@echo "Vetting code..."
	go vet ./...
	staticcheck ./...
	@echo "Running tests..."
	go test -race -vet=off ./...
.PHONY: test
test:
	@echo "Running tests..."
	go test -race ./...
	
# ======================================================= #
# BUILD 
# ======================================================= #

current_time = ${shell date --iso-8601=seconds}
git_description = ${shell git describe --always --dirty --tags --long}
linker_flags= '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build/api: build the cmd/api application
.PHONY: build/api 
build/api:
	@echo "Building cmd/api..."
	CGO_ENABLED=0 GOOS=linux go build -ldflags=${linker_flags} -o=./bin/api ./cmd/api

