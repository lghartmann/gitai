package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"huseynovvusal/gitai/internal/ai"
	"huseynovvusal/gitai/internal/git"
	"huseynovvusal/gitai/internal/ui"
)

var genCmCmd = &cobra.Command{
	Use:     "commit_message",
	Aliases: []string{"cm", "cmsg"},
	Short:   "Generate commit messages using AI",
	Long:    `The commit_message command leverages AI to generate meaningful and concise commit messages based on the changes in your repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		loaderModel := ui.NewLoaderModel()
		prog := tea.NewProgram(loaderModel)
		done := make(chan struct{})

		go func() {
			_, _ = prog.Run()
			close(done)
		}()

		diff, err := git.GetDiff()
		if err != nil {
			fmt.Println("Error getting git diff:", err)
			return
		}

		status, err := git.GetStatus()
		if err != nil {
			fmt.Println("Error getting git status:", err)
			return
		}

		commitMessage, err := ai.GenerateCommitMessage(diff, status)

		prog.Send(tea.KeyMsg{
			Type:  tea.KeyRunes,
			Runes: []rune("q"),
		})

		if err != nil {
			fmt.Println("Error generating commit message:", err)
			return
		}

		fmt.Println(commitMessage)
	},
}

func init() {
	genCmd.AddCommand(genCmCmd)
}
