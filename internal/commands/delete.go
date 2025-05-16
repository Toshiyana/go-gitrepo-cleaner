package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-github/v45/github"
	"github.com/spf13/cobra"
	"github.com/yanagimoto-toshiki/go-gitrepo-cleaner/internal/auth"
)

var forceDelete bool

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [owner/repo]",
	Short: "Delete a GitHub repository",
	Long: `Delete a GitHub repository.
This command requires the repository name in the format 'owner/repo'.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Parse owner and repo from the argument
		parts := strings.Split(args[0], "/")
		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "Error: Invalid repository format. Use 'owner/repo'\n")
			os.Exit(1)
		}
		owner := parts[0]
		repo := parts[1]

		// Get GitHub client
		client, err := auth.GetAPIClient()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Check if repository exists
		_, err = client.GetRepository(context.Background(), owner, repo)
		if err != nil {
			var ghErr *github.ErrorResponse
			if errors.As(err, &ghErr) && ghErr.Response.StatusCode == 404 {
				fmt.Fprintf(os.Stderr, "Error: Repository '%s/%s' not found\n", owner, repo)
			} else {
				fmt.Fprintf(os.Stderr, "Error checking repository: %v\n", err)
			}
			os.Exit(1)
		}

		// Confirm deletion if not forced
		if !forceDelete {
			red := color.New(color.FgRed).SprintFunc()
			fmt.Printf("Are you sure you want to delete '%s'? %s: ", args[0], red("This action cannot be undone"))
			var confirm string
			fmt.Scanln(&confirm)
			if strings.ToLower(confirm) != "y" && strings.ToLower(confirm) != "yes" {
				fmt.Println("Operation cancelled.")
				return
			}
		}

		// Delete repository
		err = client.DeleteRepository(context.Background(), owner, repo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting repository: %v\n", err)
			os.Exit(1)
		}

		// Print success message
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("%s Repository '%s' deleted successfully.\n", green("✔️"), args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Add flags
	deleteCmd.Flags().BoolVar(&forceDelete, "force", false, "Force delete without confirmation")
}
