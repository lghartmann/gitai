package cmd

import (
	"fmt"
	"huseynovvusal/gitai/internal/ai"
	"os"

	"github.com/openai/openai-go/v2/packages/param"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitai",
	Short: "GitAI is a CLI tool to interact with Git repositories using AI",
	Long:  `GitAI allows you to perform various Git operations with the help of AI, making version control easier and more intuitive.`,
}

func Execute() {
    res, err := ai.CallLLM(
        "You are a helpful AI assistant.",
        "Say hello to the user.",
        param.NewOpt[int64](256),
        param.NewOpt(0.25),
    )

    if err == nil {
        fmt.Println(res)
    } else {
        fmt.Println("Failed to call LLM:", err)
    }

    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}