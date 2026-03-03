# Makefile

# ====================================================================================
# Variables
# ====================================================================================

# Go settings
GO_PACKAGES      := ./...
GO_TEST_FLAGS    := -v -race -coverprofile=coverage.out

# ====================================================================================
# Commands (Targets)
# ====================================================================================

.PHONY: all build clean lint release test help

all: test lint build ## Run tests, lint, and build (snapshot)

help: ## Show this help message
	@echo "Usage: make <command>"
	@echo ""
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# ------------------------------------------------------------------------------------
# Development
# ------------------------------------------------------------------------------------

lint: ## Run the linter (golangci-lint)
	@echo "🔎 Running linter..."
	golangci-lint run $(GO_PACKAGES)

test: ## Run tests
	@echo "🧪 Running tests..."
	@go test $(GO_TEST_FLAGS) $(GO_PACKAGES)

# ------------------------------------------------------------------------------------
# Build and Release
# ------------------------------------------------------------------------------------

build: clean ## Generate a snapshot build with GoReleaser
	@echo "🔨 Generating build (snapshot)..."
	goreleaser build --snapshot --single-target --clean

release: clean ## Create a new release with GoReleaser
	@echo "🚀 Creating release..."
	goreleaser release --clean

# ------------------------------------------------------------------------------------
# Cleanup
# ------------------------------------------------------------------------------------

clean: ## Clean up build artifacts and caches
	@echo "🧹 Cleaning up..."
	@rm -rf ./dist
	@rm -f coverage.out
	@go clean -cache
