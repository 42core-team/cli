package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type State int

const (
	TournamentList State = iota
	TournamentDetails
)

type Model struct {
	state                 State
	tournamentListForm    *huh.Form
	tournamentDetailsForm *huh.Form
}

func NewModel() Model {
	m := Model{
		state: TournamentList,
	}
	m.tournamentListForm = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("tournamentName").
				Options(huh.NewOptions("New", "Tournament 1", "Tournament 2")...).
				Title("Choose your tournament"),
		),
	)
	m.tournamentDetailsForm = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("tournamentDetails").
				Title("Tournament Details").
				Description("Choose an option").
				Options(huh.NewOptions("Option 1", "Option 2", "Option 3")...),
		),
	)
	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.tournamentListForm.Init(), m.tournamentDetailsForm.Init())
}
