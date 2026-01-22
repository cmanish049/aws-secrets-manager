SWAG := $(shell go env GOPATH)/bin/swag

.PHONY: dev backend frontend install clean build swagger install-swag

# Run both frontend and backend concurrently
dev:
	@echo "Generating Swagger documentation..."
	@cd backend/cmd/server && $(SWAG) init -d .,../../internal/handlers,../../internal/services -o ../../docs
	@echo "Starting backend and frontend..."
	@trap 'kill 0' INT; \
	(cd backend && go run cmd/server/main.go) & \
	(cd frontend && npm run dev) & \
	wait

# Run backend only
backend:
	cd backend && go run cmd/server/main.go

# Run frontend only
frontend:
	cd frontend && npm run dev

# Install all dependencies
install:
	@echo "Installing backend dependencies..."
	cd backend && go mod download
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

# Build both projects
build:
	@echo "Building backend..."
	cd backend && go build -o bin/server cmd/server/main.go
	@echo "Building frontend..."
	cd frontend && npm run build

# Generate Swagger documentation
swagger:
	cd backend/cmd/server && $(SWAG) init -d .,../../internal/handlers,../../internal/services -o ../../docs

# Install swag CLI tool
install-swag:
	go install github.com/swaggo/swag/cmd/swag@latest

# Clean build artifacts
clean:
	rm -rf backend/bin
	rm -rf frontend/dist

# Setup project (install deps + generate swagger)
setup: install swagger
	@echo "Setup complete!"

# Help
help:
	@echo "Available commands:"
	@echo "  make dev         - Run both backend and frontend concurrently"
	@echo "  make backend     - Run backend only"
	@echo "  make frontend    - Run frontend only"
	@echo "  make install     - Install all dependencies"
	@echo "  make install-swag - Install swag CLI tool"
	@echo "  make build       - Build both projects"
	@echo "  make swagger     - Regenerate Swagger documentation"
	@echo "  make clean       - Remove build artifacts"
	@echo "  make setup       - Install dependencies and generate swagger"
	@echo "  make help        - Show this help message"
