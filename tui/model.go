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
)

type Model struct {
	state        State
	tListForm    *huh.Form
	tDetailsForm *huh.Form
	pListForm    *huh.Form
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

func initTListForm(m *Model) tea.Cmd {
	m.tListForm = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("tournamentName").
				Options(huh.NewOptions("New", "Tournament 1", "Tournament 2")...).
				Title("Tournament List").
				Description("Choose a tournament to view details"),
		),
	)
	return m.tListForm.Init()
}

func initTDetailsForm(m *Model) tea.Cmd {
	m.tDetailsForm = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("tournamentDetails").
				Title("Tournament Details").
				Description("Choose an option").
				Options(huh.NewOptions("List Participants", "Run games", "Edit Settings")...),
		),
	)
	return m.tDetailsForm.Init()
}

func initPListForm(m *Model) tea.Cmd {
	m.pListForm = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("playerName").
				Options(huh.NewOptions("New", "Player 1", "Player 2")...).
				Title("Player List").
				Description("Choose a player to view details"),
		),
	)
	return m.pListForm.Init()
}
