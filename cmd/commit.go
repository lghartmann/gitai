package cmd

import (
	"huseynovvusal/gitai/internal/tui/commit"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "AI-powered commit flow",
	Run: func(cmd *cobra.Command, args []string) {
		commit.RunCommitFlow()
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
