package commit

import (
	"huseynovvusal/gitai/internal/git"

	tea "github.com/charmbracelet/bubbletea"
)

func RunCommitFlow() {
	files, err := git.GetChangedFiles()
	if err != nil {
		panic(err)
	}
	if len(files) == 0 {
		println("No changed files to commit.")
		return
	}

	newFileSelectorModel := NewFileSelectorModel(files)

	fileSelectorProgram := tea.NewProgram(&newFileSelectorModel)
	if _, err := fileSelectorProgram.Run(); err != nil {
		panic(err)
	}

	if newFileSelectorModel.quitting {
		return
	}

	//! For demonstration, just print selected files
	// fmt.Println("Selected files for commit:")
	// for i, file := range model.files {
	// 	if model.selected[i] {
	// 		fmt.Println(" -", file)
	// 	}
	// }

	// TODO: Show error message if no files selected

	aiModel := NewAIMessageModel(newFileSelectorModel.files)
	aiModelProgram := tea.NewProgram(&aiModel)

	_, err = aiModelProgram.Run()
	if err != nil {
		panic(err)
	}

	// Print the generated commit message
	// fmt.Println(message)

}
