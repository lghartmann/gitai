package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	spinner spinner.Model
	done    bool
}

func (m *model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			m.done = true
			return m, tea.Quit
		}

	case spinner.TickMsg:
		if m.done {
			return m, tea.Quit
		}

		var cmd tea.Cmd

		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.Msg:
		if m.done {
			return m, tea.Quit
		}
	}

	return m, nil

}

func (m *model) View() string {
	if m.done {
		return "Done!\n"
	}

	return m.spinner.View() + " Loading..."
}
