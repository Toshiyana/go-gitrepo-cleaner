# Makefile for GitHub Repository Cleaner CLI

# Variables
BINARY_NAME=githubcli

# Default target
.PHONY: all
all: list

# バイナリをビルドするターゲット
.PHONY: build-cli
build-cli:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./cmd/githubcli/bin/githubcli ./cmd/githubcli/main.go

# CLI実行用の共通ターゲット
.PHONY: cli-exec
cli-exec: ## CLI実行($CMDにコマンドを指定)
	@docker compose run --rm app \
		sh -c "cd /go/src/app && go build -o ./cmd/githubcli/bin/githubcli ./cmd/githubcli/main.go && ./cmd/githubcli/bin/githubcli ${CMD}"

# 各コマンド用のターゲット
.PHONY: list
list:
	@make cli-exec CMD="list"

.PHONY: list-all
list-all:
	@make cli-exec CMD="list --all"

.PHONY: list-json
list-json:
	@make cli-exec CMD="list --json"

.PHONY: delete
delete:
	@if [ -z "$(REPO)" ]; then \
		echo "Error: REPO parameter is required. Usage: make delete REPO=owner/repo"; \
		exit 1; \
	fi
	@make cli-exec CMD="delete $(REPO)"

.PHONY: delete-force
delete-force:
	@if [ -z "$(REPO)" ]; then \
		echo "Error: REPO parameter is required. Usage: make delete-force REPO=owner/repo"; \
		exit 1; \
	fi
	@make cli-exec CMD="delete $(REPO) --force"

# ローカル開発用コマンド
.PHONY: build
build:
	@echo "Building $(BINARY_NAME) locally..."
	@go build -o $(BINARY_NAME) ./cmd/githubcli

.PHONY: install
install:
	@echo "Installing $(BINARY_NAME) locally..."
	@go install ./cmd/githubcli

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@go fmt ./...

.PHONY: lint
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed"; \
		exit 1; \
	fi

.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -rf ./cmd/githubcli/bin

# Docker関連コマンド
.PHONY: dev
dev:
	@echo "Starting development container with hot reloading..."
	@docker compose up dev

.PHONY: status
status:
	@echo "Checking Docker container status..."
	@docker compose ps

.PHONY: help
help:
	@echo "GitHub Repository Cleaner CLI - Makefile targets:"
	@echo ""
	@echo "CLI commands (Docker を使用):"
	@echo "  make list                         List repositories"
	@echo "  make list-all                     List all repositories"
	@echo "  make list-json                    List repositories in JSON format"
	@echo "  make delete REPO=owner/repo       Delete a repository with confirmation"
	@echo "  make delete-force REPO=owner/repo Delete a repository without confirmation"
	@echo ""
	@echo "Docker commands:"
	@echo "  make dev          Start development container with hot reloading"
	@echo "  make status       Check Docker container status"
	@echo ""
	@echo "Local development commands (開発用):"
	@echo "  make build        Build the binary locally"
	@echo "  make install      Install the binary locally"
	@echo "  make test         Run tests locally"
	@echo "  make test-coverage Run tests with coverage locally"
	@echo "  make fmt          Format code locally"
	@echo "  make lint         Lint code locally"
	@echo "  make clean        Clean build artifacts"
	@echo ""
