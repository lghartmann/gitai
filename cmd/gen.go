package cmd

import (
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:     "gen",
	Aliases: []string{"g", "generate"},
	Short:   "Generate Git commands or commit messages using AI",
	Long:    `The gen command leverages AI to generate Git commands or commit messages based on the current state of your repository.`,
}

func init() {
	rootCmd.AddCommand(genCmd)
}
