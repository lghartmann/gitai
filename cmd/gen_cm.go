package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"huseynovvusal/gitai/internal/ai"
	"huseynovvusal/gitai/internal/git"
	"huseynovvusal/gitai/internal/ui"
)

var (
	detailed bool
	doCommit bool
	add      bool
	push     bool
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

		commitMessage, err := ai.GenerateCommitMessage(diff, status, detailed)

		prog.Send(tea.KeyMsg{
			Type:  tea.KeyRunes,
			Runes: []rune("q"),
		})

		// Wait for the loader to finish
		<-done

		if err != nil {
			fmt.Println("Error generating commit message:", err)
			return
		}

		if add {
			err = git.AddChanges()
			if err != nil {
				fmt.Println("Error adding changes:", err)
				return
			}
			fmt.Println("Changes staged successfully.")
		}

		fmt.Println("Generated Commit Message:\n", commitMessage)

		if doCommit {
			err = git.CommitChanges(commitMessage)
			if err != nil {
				fmt.Println("Error committing changes:", err)
				return
			}

			fmt.Println("Changes committed successfully.")
		}

		if push {
			err = git.PushChanges()

			if err != nil {
				fmt.Println("Error pushing changes:", err)
				return
			}

			fmt.Println("Changes pushed successfully.")
		}

	},
}

func init() {
	genCmCmd.Flags().BoolVar(&detailed, "detailed", false, "Generate a detailed commit message")
	genCmCmd.Flags().BoolVar(&add, "add", false, "Stage all changes before committing")
	genCmCmd.Flags().BoolVar(&doCommit, "commit", false, "Commit with the generated message")
	genCmCmd.Flags().BoolVar(&push, "push", false, "Push changes after committing")

	genCmd.AddCommand(genCmCmd)
}
