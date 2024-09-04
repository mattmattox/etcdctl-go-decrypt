# Variables
GO_VERSION := 1.22
DOCKER_IMAGE := cube8021/etcdctl-go-decrypt
DOCKER_PLATFORMS := linux/amd64
DOCKER_TAG_LATEST := latest
DOCKER_TAG_VERSION := $(shell git rev-list --count HEAD)
DOCKER_TAG_SHA := $(shell git rev-parse --short HEAD)

# Default target
all: test build

# Install static analysis tools
install-analysis-tools:
	@echo "Installing static analysis tools..."
	go install golang.org/x/lint/golint@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

# Install dependency management tools
install-dependency-tools:
	@echo "Installing dependency management tools..."
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install github.com/psampaz/go-mod-outdated@latest

# Go static analysis
static-analysis: install-analysis-tools
	@echo "Running static analysis..."
	golint ./...
	staticcheck ./...
	go vet ./...

# Dependency management
dependency-management: install-dependency-tools
	@echo "Managing dependencies..."
	go mod vendor
	go mod verify
	go mod tidy

# Security scanning
security-scan: install-dependency-tools
	@echo "Running security scan..."
	gosec ./...

# Run all tests
test: static-analysis dependency-management security-scan
	@echo "All tests passed."

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker buildx build \
		--platform $(DOCKER_PLATFORMS) \
		--pull \
		-t $(DOCKER_IMAGE):v$(DOCKER_TAG_VERSION) \
		-t $(DOCKER_IMAGE):$(DOCKER_TAG_LATEST) \
		-t $(DOCKER_IMAGE):$(DOCKER_TAG_SHA) \
		--push \
		-f Dockerfile .

# Clean up
clean:
	@echo "Cleaning up..."
	rm -rf $(BACKEND_DIR)/vendor

# Build all targets
build: docker-build

# Phony targets
.PHONY: install-analysis-tools install-dependency-tools static-analysis dependency-management security-scan test docker-build clean build
