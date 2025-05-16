package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/google/go-github/v45/github"
	"github.com/spf13/cobra"
	"github.com/yanagimoto-toshiki/go-gitrepo-cleaner/internal/auth"
)

var (
	showAll    bool
	jsonOutput bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List GitHub repositories",
	Long:  `List all repositories for the authenticated GitHub user.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := auth.GetAPIClient()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Get repositories
		repos, err := client.ListRepositories(context.Background(), showAll)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching repositories: %v\n", err)
			os.Exit(1)
		}

		// Output repositories
		if jsonOutput {
			outputJSON(repos)
		} else {
			outputTable(repos)
		}
	},
}

func outputTable(repos []*github.Repository) {
	// Create a new tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)

	// Print header
	fmt.Fprintln(w, "Name\t|Private\t|Fork\t|Archived")
	fmt.Fprintln(w, "----\t|-------\t|----\t|--------")

	// Print repositories
	for _, repo := range repos {
		fmt.Fprintf(w, "%s\t|%t\t|%t\t|%t\n",
			*repo.Name,
			*repo.Private,
			*repo.Fork,
			*repo.Archived)
	}

	// Flush the tabwriter
	w.Flush()
}

func outputJSON(repos []*github.Repository) {
	// Create a slice of simplified repository information
	type RepoInfo struct {
		Name     string `json:"name"`
		Private  bool   `json:"private"`
		Fork     bool   `json:"fork"`
		Archived bool   `json:"archived"`
	}

	var repoInfos []RepoInfo
	for _, repo := range repos {
		repoInfos = append(repoInfos, RepoInfo{
			Name:     *repo.Name,
			Private:  *repo.Private,
			Fork:     *repo.Fork,
			Archived: *repo.Archived,
		})
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(repoInfos, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling to JSON: %v\n", err)
		os.Exit(1)
	}

	// Print JSON
	fmt.Println(string(jsonData))
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Add flags
	listCmd.Flags().BoolVar(&showAll, "all", false, "Show all repositories including archived and forked")
	listCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output in JSON format")
}
