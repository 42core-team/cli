package tui

import (
	"core-cli/db"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func initTListForm(m *Model) tea.Cmd {
	var nameList []string = []string{"<New>"}

	for _, team := range db.GetTeams() {
		nameList = append(nameList, team.Name)
	}

	m.tListForm = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("teamName").
				Options(huh.NewOptions(nameList...)...).
				Title("Team List").
				Description("Choose a team to view details or create a new one"),
		),
	)
	return m.tListForm.Init()
}

func updateTListForm(m *Model, msg *tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	form, cmd := m.tListForm.Update(*msg)
	if f, ok := form.(*huh.Form); ok {
		m.tListForm = f
		cmds = append(cmds, cmd)
	}

	if m.tListForm.State == huh.StateCompleted {
		m.mcontext.CurrentTeamName = m.tListForm.GetString("teamName")
		return switchState(m, TDetailsState)
	}

	return m, tea.Batch(cmds...)
}

func initTDetailsForm(m *Model) tea.Cmd {
	var nameList []string = []string{"<New>"}

	for _, player := range db.GetPLayersByTeamName(m.mcontext.CurrentTeamName) {
		nameList = append(nameList, player.IntraName)
	}

	m.mcontext.CurrentTeamID = db.GetTeamByName(m.mcontext.CurrentTeamName).ID

	m.tDetailsForm = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("teamDetails").
				Title("Team Details").
				Description("Choose an option").
				Options(huh.NewOptions(nameList...)...),
		),
	)
	return m.tDetailsForm.Init()
}

func updateTDetailsForm(m *Model, msg *tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	form, cmd := m.tDetailsForm.Update(*msg)
	if f, ok := form.(*huh.Form); ok {
		m.tDetailsForm = f
		cmds = append(cmds, cmd)
	}

	if m.tDetailsForm.State == huh.StateCompleted {
		switch m.tDetailsForm.GetString("teamDetails") {
		case "<New>":
			return switchState(m, PAddState)
		default:
			m.mcontext.CurrentGithubName = m.tDetailsForm.GetString("teamDetails")
			m.mcontext.CurrentPlayerID = db.GetPlayerByIntraName(m.mcontext.CurrentGithubName).ID
			return switchState(m, PDetailsState)
		}
	}

	return m, tea.Batch(cmds...)
}
