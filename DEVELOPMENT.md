# Development Guide for GitHub Repository Cleaner CLI

This document provides instructions for developing and testing the GitHub Repository Cleaner CLI tool.

## Prerequisites

- Go 1.24 or higher
- Docker and Docker Compose (optional, for containerized development)
- Docker Compose V2 (uses `docker compose` command format)
- GitHub Personal Access Token with `repo` and `delete_repo` permissions

## Setting Up the Development Environment

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-gitrepo-cleaner.git
cd go-gitrepo-cleaner
```

### 2. Set Up Authentication

Create a `.env` file in the project root:

```bash
cp .env.example .env
```

Edit the `.env` file and add your GitHub Personal Access Token:

```
GITHUB_TOKEN=your_github_token_here
```

## Development Methods

### Method 1: Direct Development

#### Using Makefile

The project includes a Makefile with common commands for building, running, and testing the CLI tool:

```bash
# Show available commands
make help

# Local development commands (開発用)
make build        # Build the application locally
make install      # Install the binary locally
make test         # Run tests locally
make fmt          # Format code locally
make lint         # Lint code locally
make clean        # Clean build artifacts

# CLI commands (Docker を使用)
# These commands will build the binary inside a Docker container and execute it
make list                         # List repositories
make list-all                     # List all repositories
make list-json                    # List repositories in JSON format
make delete REPO=owner/repo       # Delete a repository with confirmation
make delete-force REPO=owner/repo # Delete a repository without confirmation
```

#### Building the Application Manually

```bash
# Build the application
go build -o githubcli ./cmd/githubcli

# Run the application
./githubcli --help
```

#### Hot Reloading with Air

For automatic rebuilding when code changes:

```bash
# Install Air if you don't have it
go install github.com/cosmtrek/air@latest

# Run with Air for hot reloading
air

# Or using the Makefile
make dev
```

### Method 2: Docker Development

#### Using Docker for CLI Commands

The Makefile provides targets that use Docker to build and run the CLI commands:

```bash
# Run CLI commands using Docker
make list
make list-all
make list-json
make delete REPO=owner/repo
make delete-force REPO=owner/repo
```

Under the hood, these commands:
1. Create a temporary Docker container
2. Build the binary inside the container
3. Execute the command
4. Remove the container when done

This approach ensures a clean environment for each command execution.

#### Development Container with Hot Reloading (Air)

The Docker Compose configuration includes a `dev` service that automatically sets up Air for hot reloading:

```bash
# Start the development container with Air for hot reloading
make dev

# Wait for the container to fully start and Air to initialize
# You should see output indicating that Air is watching for changes

# Air will automatically rebuild and restart the application when code changes are detected
# The binary is available at /go/src/app/tmp/main
```

This setup includes:
- Automatic installation of Air
- Volume mapping for code changes
- Persistent Go module cache for faster builds
- Automatic rebuilding when files change

## Testing the Functionality

### 1. Testing the List Command

```bash
# Test listing repositories (non-archived, non-forked)
./githubcli list

# Test listing all repositories
./githubcli list --all

# Test JSON output
./githubcli list --json
```

With Docker:

```bash
make list
make list-all
make list-json
```

### 2. Testing the Delete Command

**CAUTION**: This will actually delete repositories. Use with care.

```bash
# Test deleting a repository (with confirmation prompt)
./githubcli delete yourusername/test-repo

# Test force delete (no confirmation)
./githubcli delete yourusername/test-repo --force
```

With Docker:

```bash
make delete REPO=yourusername/test-repo
make delete-force REPO=yourusername/test-repo
```

### 3. Creating Test Repositories for Safe Testing

You can create temporary test repositories on GitHub for testing the delete functionality:

1. Go to GitHub and create a new repository (e.g., "test-repo-for-deletion")
2. Clone it locally
3. Test the delete command on this repository

## Debugging

### Enabling Verbose Output

You can add a verbose flag to the commands for debugging:

```go
// Add to root.go
var verbose bool

func init() {
    rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}
```

### Using Delve Debugger

```bash
# Install Delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug the application
dlv debug ./cmd/githubcli/main.go
```

## Common Issues and Solutions

### Authentication Issues

If you encounter authentication errors:

1. Verify your token has the correct permissions (`repo` and `delete_repo`)
2. Check that the token is correctly set in the `.env` file or as an environment variable
3. Try regenerating your GitHub token

### Rate Limiting

GitHub API has rate limits. If you hit them:

1. Wait for the rate limit to reset
2. Use authenticated requests (which have higher rate limits)
3. Implement exponential backoff in the code for retries

## Continuous Integration

You can set up GitHub Actions for CI/CD:

1. Create a `.github/workflows/go.yml` file
2. Configure it to run tests and build the application
3. Use GitHub Secrets to store your GitHub token for testing
