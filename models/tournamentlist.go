package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type TournamentListModel struct {
	main *MainModel
	form *huh.Form
}

func NewTournamentListModel(main *MainModel) TournamentListModel {
	m := TournamentListModel{
		main: main,
	}
	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("tournamentName").
				Options(huh.NewOptions("New", "Tournament 1", "Tournament 2")...).
				Title("Choose your tournament"),
		),
	)
	return m
}

func (m TournamentListModel) Init() tea.Cmd {
	return tea.Batch(m.form.Init())
}

func (m TournamentListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	var cmds []tea.Cmd
	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		m.main.currentModelType = ModelTournament
		m.main.lastModel = modelTournamentList
		m.main.currentModel = NewTournamentModel(m.main)
		cmds = append(cmds, m.main.currentModel.Init())
		return m, tea.Batch(cmds...)
	}

	return m, tea.Batch(cmds...)
}

func (m TournamentListModel) View() string {
	return m.form.View()
}
