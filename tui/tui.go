package tui

import (
	tea "github.com/charmbracelet/bubbletea"
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
			case PAddState:
				return switchState(&m, TDetailsState)
			case PDetailsState:
				return switchState(&m, TDetailsState)
			}
		}
	}

	switch m.state {
	case TListState:
		return updateTListForm(&m, &msg)
	case TDetailsState:
		return updateTDetailsForm(&m, &msg)
	case PAddState:
		return updatePAddForm(&m, &msg)
	case PDetailsState:
		return updatePDetailsForm(&m, &msg)
	}

	return m, nil
}

func (m Model) View() string {
	switch m.state {
	case TListState:
		return m.tListForm.View()
	case TDetailsState:
		return m.tDetailsForm.View()
	case PAddState:
		return m.pAddForm.View()
	case PDetailsState:
		return m.pDetailsForm.View()
	}
	return "Empty view"
}
