package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type State int

const (
	TListState State = iota
	TDetailsState
	PAddState
	PDetailsState
)

type Model struct {
	state        State
	mcontext     ModelContext
	tListForm    *huh.Form
	tDetailsForm *huh.Form
	pAddForm     *huh.Form
	pDetailsForm *huh.Form
}

type ModelContext struct {
	CurrentTeamID     uint
	CurrentTeamName   string
	CurrentPlayerID   uint
	CurrentGithubName string
	CurrentIntraName  string
}

func NewModel() Model {
	m := Model{
		state: TListState,
		mcontext: ModelContext{
			CurrentTeamName:   "",
			CurrentTeamID:     0,
			CurrentGithubName: "",
			CurrentPlayerID:   0,
		},
	}
	initTListForm(&m)
	return m
}

func (m Model) Init() tea.Cmd {
	return initTListForm(&m)
}
