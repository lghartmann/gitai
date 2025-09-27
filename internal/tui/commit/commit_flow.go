package commit

import (
	"fmt"
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

	model := NewFileSelectorModel(files)

	p := tea.NewProgram(&model)

	if _, err := p.Run(); err != nil {
		panic(err)
	}

	//! For demonstration, just print selected files
	fmt.Println("Selected files for commit:")
	for i, file := range model.files {
		if model.selected[i] {
			fmt.Println(" -", file)
		}
	}
}
