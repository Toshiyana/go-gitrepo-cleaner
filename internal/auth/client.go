package auth

import (
	"github.com/yanagimoto-toshiki/go-gitrepo-cleaner/internal/api"
)

// GetAPIClient returns an authenticated GitHub API client
func GetAPIClient() (*api.GitHubClient, error) {
	client, err := GetGitHubClient()
	if err != nil {
		return nil, err
	}
	return api.NewGitHubClient(client), nil
}
