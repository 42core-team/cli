package tui

import tea "github.com/charmbracelet/bubbletea"

func switchState(m *Model, state State) (tea.Model, tea.Cmd) {
	m.state = state
	switch state {
	case TListState:
		return *m, initTListForm(m)
	case TDetailsState:
		return *m, initTDetailsForm(m)
	case PAddState:
		return *m, initPAddForm(m)
	case PDetailsState:
		return *m, initPDetailsForm(m)
	default:
		return *m, nil
	}
}
