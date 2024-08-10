package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type State int

const (
	TListState State = iota
	TDetailsState
	PListState
	PAddState
)

type Model struct {
	state        State
	tListForm    *huh.Form
	tDetailsForm *huh.Form
	pListForm    *huh.Form
	pAddForm     *huh.Form
}

func NewModel() Model {
	m := Model{
		state: TListState,
	}
	initTListForm(&m)
	return m
}

func (m Model) Init() tea.Cmd {
	return initTListForm(&m)
}
