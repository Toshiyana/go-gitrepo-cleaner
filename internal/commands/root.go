package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "githubcli",
	Short: "GitHub Repository Management CLI",
	Long: `A CLI tool for managing GitHub repositories.
This tool allows you to list and delete repositories from your GitHub account.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Add any global flags here
}
