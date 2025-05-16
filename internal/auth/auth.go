package auth

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/go-github/v45/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

// GetGitHubClient returns an authenticated GitHub client
func GetGitHubClient() (*github.Client, error) {
	// Try to load .env file from current directory
	_ = godotenv.Load()

	// Get GitHub token from environment variable
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		// Try to load from home directory
		home, err := os.UserHomeDir()
		if err == nil {
			_ = godotenv.Load(filepath.Join(home, ".env"))
			token = os.Getenv("GITHUB_TOKEN")
		}

		if token == "" {
			return nil, errors.New("GitHub token not found. Please set GITHUB_TOKEN environment variable or add it to .env file")
		}
	}

	// Create OAuth2 token source
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	// Create GitHub client
	client := github.NewClient(tc)

	// Verify token by getting authenticated user
	_, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate with GitHub: %v", err)
	}

	return client, nil
}
