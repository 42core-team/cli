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
	mcontext     ModelContext
	tListForm    *huh.Form
	tDetailsForm *huh.Form
	pListForm    *huh.Form
	pAddForm     *huh.Form
}

type ModelContext struct {
	CurrentTeamName   string
	CurrentTeamID     int
	CurrentPlayerName string
	CurrentPlayerID   int
}

func NewModel() Model {
	m := Model{
		state: TListState,
		mcontext: ModelContext{
			CurrentTeamName:   "",
			CurrentTeamID:     0,
			CurrentPlayerName: "",
			CurrentPlayerID:   0,
		},
	}
	initTListForm(&m)
	return m
}

func (m Model) Init() tea.Cmd {
	return initTListForm(&m)
}
