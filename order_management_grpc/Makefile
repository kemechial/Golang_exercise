# Order Management gRPC Project Makefile

.PHONY: proto build server client test clean install-deps

# Generate proto files
proto:
	@echo "Generating proto files..."
	@./scripts/generate-proto.sh

# Install dependencies
install-deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod download

# Build all binaries
build: proto
	@echo "Building server..."
	@go build -o bin/server cmd/server/main.go
	@echo "Building client..."
	@go build -o bin/client cmd/client/main.go

# Run server
server: build
	@echo "Starting Order Management server..."
	@./bin/server

# Run client
client: build
	@echo "Running Order Management client..."
	@./bin/client

# Run tests
test:
	@echo "Running unit and integration tests..."
	@go test -v ./...

test-integration:
	@echo "Running integration tests in Docker..."
	@docker compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from test
	@docker compose -f docker-compose.test.yml down

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -f proto/*.pb.go

# Install protoc plugins
install-protoc:
	@echo "Installing protoc plugins..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	@golangci-lint run

# Run development server with auto-reload
dev:
	@echo "Starting development server..."
	@air -c .air.toml

# Docker build
docker-build:
	@echo "Building Docker image..."
	@docker build -t order-management:latest .

# Help
help:
	@echo "Available commands:"
	@echo "  proto         - Generate proto files"
	@echo "  install-deps  - Install Go dependencies"
	@echo "  build         - Build server and client"
	@echo "  server        - Run the gRPC server"
	@echo "  client        - Run the test client"
	@echo "  test          - Run tests"
	@echo "  clean         - Clean build artifacts"
	@echo "  install-protoc- Install protoc plugins"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code"
	@echo "  docker-build  - Build Docker image"
	@echo "  help          - Show this help"
