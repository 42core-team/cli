package tui

import tea "github.com/charmbracelet/bubbletea"

func switchState(m *Model, state State) (tea.Model, tea.Cmd) {
	m.state = state
	switch state {
	case TListState:
		return *m, initTListForm(m)
	case TDetailsState:
		return *m, initTDetailsForm(m)
	case PListState:
		return *m, initPListForm(m)
	default:
		return *m, nil
	}
}
