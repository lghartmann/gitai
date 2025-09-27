package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitai",
	Short: "Gitai is a CLI tool to interact with Git repositories using AI",
	Long:  `Gitai allows you to perform various Git operations with the help of AI, making version control easier and more intuitive.`,
}

func Execute() {

	password := "commit test"
	fmt.Println(password)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
