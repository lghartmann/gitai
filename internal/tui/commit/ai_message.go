package commit

import (
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"

	"huseynovvusal/gitai/internal/ai"
	"huseynovvusal/gitai/internal/git"
	"huseynovvusal/gitai/internal/tui/commit/shared"
)

type AIMessageModel struct {
	files   []string
	message string
	done    bool
	cancel  bool
	spinner spinner.Model
}

type aiDoneMsg struct {
	message string
}

func NewAIMessageModel(files []string) AIMessageModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = shared.CursorStyle

	return AIMessageModel{
		files:   files,
		message: "",
		done:    false,
		cancel:  false,
		spinner: s,
	}
}

func runAIAsync(files []string) tea.Cmd {
	return func() tea.Msg {
		diff, err := git.GetChangesForFiles(files)
		if err != nil {
			panic(err)
		}

		status, err := git.GetStatus()
		if err != nil {
			panic(err)
		}

		// TODO: For now, we assume user always wants to use concide commit message
		commitMessage, err := ai.GenerateCommitMessage(diff, status, false)

		if err != nil {
			panic(err)
		}

		return aiDoneMsg{message: commitMessage}
	}
}

func (m *AIMessageModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		runAIAsync(m.files),
	)
}

func (m *AIMessageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "e":
			// TODO: Open editor to edit the commit message
		case "r":
			// TODO: Regenerate the commit message
		case "c":
			// TODO: Commit the changes with the AI-generated message
		case "x":
			m.cancel = true
			return m, tea.Quit
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case aiDoneMsg:
		m.done = true
		m.message = msg.message
		return m, nil
	}

	return m, nil
}

func (m *AIMessageModel) View() string {
	if m.cancel {
		return shared.ErrorStyle.Render("Commit cancelled.") + "\n"
	}

	if !m.done {
		return "\n" + shared.HeaderStyle.Render("Generating commit message...") + "\n\n" + m.spinner.View() + " Generating commit message..." + "\n"
	}

	var b strings.Builder

	header := shared.HeaderStyle.Render("AI commit message suggestion:")
	b.WriteString("\n" + header + "\n")
	b.WriteString(m.message + "\n")

	b.WriteString("\n[y] Approve   [e] Edit   [r] Regenerate   [c] Commit   [x] Cancel\n")

	return b.String()

}
