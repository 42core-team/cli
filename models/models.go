package models

import tea "github.com/charmbracelet/bubbletea"

type ModelType int

const (
	modelNil ModelType = iota
	modelTournamentList
	ModelTournament
)

func GetModelConstructor(model ModelType, mainModel *MainModel) tea.Model {
	switch model {
	case modelTournamentList:
		return NewTournamentListModel(mainModel)
	case ModelTournament:
		return NewTournamentModel(mainModel)
	default:
		return nil
	}
}
