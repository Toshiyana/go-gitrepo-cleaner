# GitHub Repository Cleaner CLI

A command-line tool for managing GitHub repositories. This tool allows you to list and delete repositories from your GitHub account.

## Features

- List all repositories (public and private)
- Filter repositories by type (archived, forked)
- Delete repositories with confirmation prompt
- JSON output option for scripting

## Installation

### Prerequisites

- Go 1.24 or higher
- GitHub Personal Access Token with `repo` and `delete_repo` permissions

### Install from source

```bash
go install github.com/yanagimoto-toshiki/go-gitrepo-cleaner/cmd/githubcli@latest
```

## Authentication

The tool requires a GitHub Personal Access Token (PAT) with appropriate permissions. You can provide this token in one of two ways:

1. Set the `GITHUB_TOKEN` environment variable:

```bash
export GITHUB_TOKEN=your_github_token
```

2. Create a `.env` file in your home directory or the current directory:

```
GITHUB_TOKEN=your_github_token
```

## Usage

### List Repositories

```bash
# List non-archived, non-forked repositories
githubcli list

# List all repositories including archived and forked
githubcli list --all

# Output in JSON format
githubcli list --json
```

### Delete Repository

```bash
# Delete a repository with confirmation prompt
githubcli delete owner/repo

# Delete a repository without confirmation
githubcli delete owner/repo --force
```

## Development

### Setup

1. Clone the repository:

```bash
git clone https://github.com/yourusername/go-gitrepo-cleaner.git
cd go-gitrepo-cleaner
```

2. Install dependencies:

```bash
go mod download
```

3. Build the application:

```bash
go build -o githubcli ./cmd/githubcli
```

### Using Docker for Development

A Docker environment is provided for development and testing:

```bash
# Build and run using Docker Compose
docker-compose up -d

# Execute commands inside the container
docker-compose exec app githubcli list
```

## License

MIT
