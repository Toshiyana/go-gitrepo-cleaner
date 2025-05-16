package api

import (
	"context"

	"github.com/google/go-github/v45/github"
)

// GitHubClient wraps the GitHub API client
type GitHubClient struct {
	client *github.Client
}

// NewGitHubClient creates a new GitHubClient
func NewGitHubClient(client *github.Client) *GitHubClient {
	return &GitHubClient{
		client: client,
	}
}

// ListRepositories lists all repositories for the authenticated user
func (c *GitHubClient) ListRepositories(ctx context.Context, showAll bool) ([]*github.Repository, error) {
	// Set up options for listing repositories
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	// If showAll is false, exclude archived and forked repositories
	if !showAll {
		opt.Affiliation = "owner"
	}

	// Get all repositories for the authenticated user
	var allRepos []*github.Repository
	for {
		repos, resp, err := c.client.Repositories.List(ctx, "", opt)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	// Filter out archived and forked repositories if showAll is false
	if !showAll {
		var filteredRepos []*github.Repository
		for _, repo := range allRepos {
			if repo != nil && !*repo.Archived && !*repo.Fork {
				filteredRepos = append(filteredRepos, repo)
			}
		}
		return filteredRepos, nil
	}

	return allRepos, nil
}

// GetRepository gets a repository by owner and name
func (c *GitHubClient) GetRepository(ctx context.Context, owner, repo string) (*github.Repository, error) {
	repository, _, err := c.client.Repositories.Get(ctx, owner, repo)
	return repository, err
}

// DeleteRepository deletes a repository
func (c *GitHubClient) DeleteRepository(ctx context.Context, owner, repo string) error {
	_, err := c.client.Repositories.Delete(ctx, owner, repo)
	return err
}
