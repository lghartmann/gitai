package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
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
		diff, err := git.GetDiff()
		if err != nil {
			fmt.Println("❌ Error getting git diff:", err)
			return
		}

		status, err := git.GetStatus()
		if err != nil {
			fmt.Println("❌ Error getting git status:", err)
			return
		}

		// Check for sensitive data before starting any loaders
		sensitiveData, err := ai.CheckDiffSafety(diff)
		if err != nil {
			fmt.Println("❌ Error checking diff safety:", err)
			return
		}

		if len(sensitiveData) > 0 {
			// Show warnings
			fmt.Println("⚠️  WARNING: Potential sensitive data detected:")
			for _, data := range sensitiveData {
				fmt.Printf("  - %s\n", data)
			}

			proceed := false
			prompt := &survey.Confirm{
				Message: "Do you want to proceed with the commit?",
				Default: false,
			}
			err := survey.AskOne(prompt, &proceed)
			if err != nil {
				fmt.Println("❌ Error reading input:", err)
				return
			}

			if !proceed {
				fmt.Println("Commit generation canceled.")
				return
			}
		}

		// Only start the loader for the AI call (the slow part)
		loaderModel := ui.NewLoaderModel()
		prog := tea.NewProgram(loaderModel)
		done := make(chan struct{})

		go func() {
			_, _ = prog.Run()
			close(done)
		}()

		commitMessage, err := ai.GenerateCommitMessage(diff, status, detailed)

		// Stop the loader
		prog.Send(tea.KeyMsg{
			Type:  tea.KeyRunes,
			Runes: []rune("q"),
		})
		<-done

		if err != nil {
			fmt.Println("❌ Error generating commit message:", err)
			return
		}

		if add {
			err = git.AddChanges()
			if err != nil {
				fmt.Println("❌ Error adding changes:", err)
				return
			}
			fmt.Println("✅ Changes staged successfully.")
		}

		fmt.Println("📝 Generated Commit Message:\n", commitMessage)

		if doCommit {
			err = git.CommitChanges(commitMessage)
			if err != nil {
				fmt.Println("❌ Error committing changes:", err)
				return
			}

			fmt.Println("✅ Changes committed successfully.")
		}

		if push {
			err = git.PushChanges()

			if err != nil {
				fmt.Println("❌ Error pushing changes:", err)
				return
			}

			fmt.Println("🚀 Changes pushed successfully.")
		}

	},
}

func init() {
	genCmCmd.Flags().BoolVarP(&detailed, "detailed", "d", false, "Generate a detailed commit message")
	genCmCmd.Flags().BoolVarP(&add, "add", "a", false, "Stage all changes before committing")
	genCmCmd.Flags().BoolVarP(&doCommit, "commit", "c", false, "Commit with the generated message")
	genCmCmd.Flags().BoolVarP(&push, "push", "p", false, "Push changes after committing")

	genCmd.AddCommand(genCmCmd)
}
