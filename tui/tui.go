package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd
	switch m.state {
	case TournamentList:
		form, cmd := m.tournamentListForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.tournamentListForm = f
			cmds = append(cmds, cmd)
		}
		if m.tournamentListForm.State == huh.StateCompleted {
			m.state = TournamentDetails
		}
	case TournamentDetails:
		form, cmd := m.tournamentDetailsForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.tournamentDetailsForm = f
			cmds = append(cmds, cmd)
		}
		if m.tournamentDetailsForm.State == huh.StateCompleted {
			m.state = TournamentList
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	switch m.state {
	case TournamentList:
		return m.tournamentListForm.View()
	case TournamentDetails:
		return m.tournamentDetailsForm.View()
	}
	return "Empty view"
}
