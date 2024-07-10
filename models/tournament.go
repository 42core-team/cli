package models

import (
	"core-cli/data"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type TournamentModel struct {
	data *data.Data
	form *huh.Form
}

func NewTournamentModel(data.Data *data.Data) TournamentModel {
	m := TournamentModel{
		main: main,
	}
	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("tournamentOption").
				Options(huh.NewOptions("List Participants", "Run Game(s)", "Edit Settings")...).
				Title("Choose an option"),
		),
	)
	return m
}

func (m TournamentModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m TournamentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m TournamentModel) View() string {
	if m.form.State == huh.StateCompleted {
		tournament := m.form.GetString("tournamentOption")
		return fmt.Sprintf("You selected: %s", tournament)
	}
	return m.form.View()
}
