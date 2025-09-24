package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var genCmCmd = &cobra.Command{
	Use:     "commit_message",
	Aliases: []string{"cm", "cmsg"},
	Short:   "Generate commit messages using AI",
	Long:    `The commit_message command leverages AI to generate meaningful and concise commit messages based on the changes in your repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation for generating commit messages will go here
		fmt.Println("Generating commit message...")
	},
}

func init() {
	genCmd.AddCommand(genCmCmd)
}
