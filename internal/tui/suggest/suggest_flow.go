package suggest

import (
	"huseynovvusal/gitai/internal/git"

	tea "github.com/charmbracelet/bubbletea"
)

func RunSuggestFlow() {
	files, err := git.GetChangedFiles()
	if err != nil {
		panic(err)
	}

	if len(files) == 0 {
		println("No changed files to commit.")
		return
	}

	fileSelectorModel := NewFileSelectorModel(files)
	fileSelectorProgram := tea.NewProgram(&fileSelectorModel)
	if _, err := fileSelectorProgram.Run(); err != nil {
		panic(err)
	}

	if fileSelectorModel.quitting {
		return
	}

	// TODO: Show error message if no files selected

	if len(fileSelectorModel.files) == 0 {
		println("No files selected.")
		return
	}

	aiModel := NewAIMessageModel(fileSelectorModel.files)
	aiModelProgram := tea.NewProgram(&aiModel)

	_, err = aiModelProgram.Run()
	if err != nil {
		panic(err)
	}

}
