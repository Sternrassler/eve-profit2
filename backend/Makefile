# Makefile for EVE Profit Calculator 2.0 Backend

.PHONY: help test test-unit test-integration test-all coverage clean build run deps

# Default target
help:
	@echo "Available targets:"
	@echo "  deps           - Install dependencies"
	@echo "  build          - Build the application"
	@echo "  run            - Run the application"
	@echo "  test           - Run all tests"
	@echo "  test-unit      - Run unit tests only"
	@echo "  test-integration - Run integration tests only"
	@echo "  coverage       - Generate test coverage report"
	@echo "  clean          - Clean build artifacts"
	@echo "  lint           - Run linting tools"

# Install dependencies
deps:
	go mod download
	go mod tidy

# Build the application
build:
	go build -o bin/eve-profit2 ./cmd/server

# Run the application
run:
	go run ./cmd/server

# Run all tests
test: test-unit test-integration

# Run unit tests only
test-unit:
	go test ./tests/unit/... -v

# Run integration tests only (requires SDE database)
test-integration:
	@if [ ! -f "data/sqlite-latest.sqlite" ]; then \
		echo "SDE database not found. Please run: bash scripts/download-sde.sh"; \
		exit 1; \
	fi
	go test ./tests/integration/... -v

# Run all tests and generate coverage
coverage:
	@if [ ! -f "data/sqlite-latest.sqlite" ]; then \
		echo "SDE database not found. Please run: bash scripts/download-sde.sh"; \
		exit 1; \
	fi
	go test -coverpkg=./internal/... -coverprofile=coverage.out ./tests/unit/... ./tests/integration/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	rm -f bin/eve-profit2
	rm -f coverage.out
	rm -f coverage.html

# Run linting tools
lint:
	@if command -v golangci-lint >/dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run; \
	fi

# Setup development environment
dev-setup: deps
	@echo "Setting up development environment..."
	@if [ ! -f "data/sqlite-latest.sqlite" ]; then \
		echo "Downloading SDE database..."; \
		bash scripts/download-sde.sh; \
	fi
	@echo "Development environment ready!"

# Run tests in CI mode (short tests only)
test-ci:
	go test -short ./tests/unit/... -v

# Quick verification before commit
pre-commit: lint test-unit
	@echo "Pre-commit checks passed!"
