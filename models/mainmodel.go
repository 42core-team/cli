package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	currentModel     tea.Model
	currentModelType ModelType
	lastModel        ModelType
}

func NewMainModel() MainModel {
	m := MainModel{
		lastModel: modelNil,
	}
	m.currentModelType = modelTournamentList
	m.currentModel = NewTournamentListModel(&m)
	return m
}

func (m MainModel) Init() tea.Cmd {
	return m.currentModel.Init()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc", "backspace":
			if m.lastModel == modelNil {
				return m, tea.Quit
			}
			newLastModelType := m.currentModelType
			m.currentModel = GetModelConstructor(m.lastModel, &m)
			m.lastModel = newLastModelType
			cmds = append(cmds, m.currentModel.Init())
			return m, tea.Batch(cmds...)
		}
	}

	_, modelCmcmds := m.currentModel.Update(msg)
	cmds = append(cmds, modelCmcmds)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	return m.currentModel.View()
}
