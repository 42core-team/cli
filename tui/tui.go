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
	case TListState:
		form, cmd := m.tListForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.tListForm = f
			cmds = append(cmds, cmd)
		}
		if m.tListForm.State == huh.StateCompleted {
			m.state = TDetailsState
			cmds = append(cmds, initTDetailsForm(&m))
		}
	case TDetailsState:
		form, cmd := m.tDetailsForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.tDetailsForm = f
			cmds = append(cmds, cmd)
		}
		if m.tDetailsForm.State == huh.StateCompleted {
			m.state = TListState
			cmds = append(cmds, initTListForm(&m))
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
	}
	return "Empty view"
}
