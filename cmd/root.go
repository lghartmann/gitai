package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitai",
	Short: "GitAI is a CLI tool to interact with Git repositories using AI",
	Long:  `GitAI allows you to perform various Git operations with the help of AI, making version control easier and more intuitive.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default action when no subcommands are provided
		fmt.Println("Welcome to GitAI! Use --help to see available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}