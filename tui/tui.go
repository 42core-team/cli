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
		case "esc":
			switch m.state {
			case TDetailsState:
				return switchState(&m, TListState)
			case PListState:
				return switchState(&m, TDetailsState)
			}
		}
	}

	var cmds []tea.Cmd
	switch m.state {
	case TListState:
		form, cmd := m.tListForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.tListForm = f
			cmds = append(cmds, cmd)
		}
		if m.tListForm.State == huh.StateCompleted {
			return switchState(&m, TDetailsState)
		}
	case TDetailsState:
		form, cmd := m.tDetailsForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.tDetailsForm = f
			cmds = append(cmds, cmd)
		}
		if m.tDetailsForm.State == huh.StateCompleted {
			switch m.tDetailsForm.GetString("tournamentDetails") {
			case "List Participants":
				return switchState(&m, PListState)
			default:
				return switchState(&m, TListState)
			}
		}
	case PListState:
		form, cmd := m.pListForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.pListForm = f
			cmds = append(cmds, cmd)
		}
		if m.pListForm.State == huh.StateCompleted {
			return switchState(&m, TDetailsState)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	switch m.state {
	case TListState:
		return m.tListForm.View()
	case TDetailsState:
		return m.tDetailsForm.View()
	case PListState:
		return m.pListForm.View()
	}
	return "Empty view"
}
