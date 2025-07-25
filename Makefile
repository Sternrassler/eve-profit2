# EVE Profit Calculator 2.0 - Makefile
# Alternative zu dev-server.sh für Make-User

.PHONY: help start stop restart status backend frontend logs test clean

# Default target
help: ## Show this help message
	@echo "EVE Profit Calculator 2.0 - Development Commands"
	@echo ""
	@echo "Usage: make [TARGET]"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-12s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "URLs:"
	@echo "  Backend:  http://localhost:9000"
	@echo "  Frontend: http://localhost:3001"

start: ## Start both backend and frontend servers
	@./dev-server.sh start

stop: ## Stop all running servers
	@./dev-server.sh stop

restart: ## Restart all servers
	@./dev-server.sh restart

status: ## Show server status
	@./dev-server.sh status

backend: ## Start only backend server
	@./dev-server.sh backend

frontend: ## Start only frontend server
	@./dev-server.sh frontend

logs: ## Show logs from both servers
	@./dev-server.sh logs

logs-backend: ## Show backend logs only
	@./dev-server.sh logs b

logs-frontend: ## Show frontend logs only
	@./dev-server.sh logs f

test: ## Run backend tests
	@./dev-server.sh test

test-e2e: ## Run E2E tests with Playwright
	@echo "Running E2E tests..."
	@npx playwright test

clean: ## Clean up log files and stop all servers
	@./dev-server.sh stop
	@rm -f backend.log frontend.log
	@echo "Clean up completed"

install: ## Install all dependencies
	@echo "Installing Go dependencies..."
	@cd backend && go mod tidy
	@echo "Installing Node.js dependencies..."
	@cd frontend && npm install
	@echo "Dependencies installed successfully"

dev: start ## Alias for start (common convention)

build: ## Build both backend and frontend for production
	@echo "Building backend..."
	@cd backend && go build -o bin/eve-profit2 cmd/server/main.go
	@echo "Building frontend..."
	@cd frontend && npm run build
	@echo "Build completed successfully"

docker-build: ## Build Docker images (if Dockerfiles exist)
	@echo "Docker build not yet implemented"

deploy: ## Deploy to production (placeholder)
	@echo "Deployment not yet implemented"
