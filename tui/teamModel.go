package tui

import (
	"core-cli/db"
	"core-cli/github"
	"core-cli/model"
	"errors"

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
			return switchState(m, TListState)
		}
	}

	return m, tea.Batch(cmds...)
}

func initPAddForm(m *Model) tea.Cmd {
	m.pAddForm = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("githubName").
				Title("Add Player").
				Description("Enter the github username of the player").
				Validate(func(input string) error {
					if input == "" {
						return errors.New("player name cannot be empty")
					}
					if db.PlayerExistsByGithubName(input) {
						return errors.New(input + " already exists in the db")
					}
					if !github.GithubUserExists(input) {
						return errors.New(input + " does not exist on github")
					}
					return nil
				}),
			huh.NewInput().
				Key("intraName").
				Description("Enter the intra username of the player").
				Validate(func(input string) error {
					if input == "" {
						return errors.New("player name cannot be empty")
					}
					if db.PlayerExistsByIntraName(input) {
						return errors.New(input + " already exists in the db")
					}
					return nil
				}),
		),
	)
	return m.pAddForm.Init()
}

func updatePAddForm(m *Model, msg *tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	form, cmd := m.pAddForm.Update(*msg)
	if f, ok := form.(*huh.Form); ok {
		m.pAddForm = f
		cmds = append(cmds, cmd)
	}

	if m.pAddForm.State == huh.StateCompleted {
		db.SavePlayer(&model.Player{
			GithubName: m.pAddForm.GetString("githubName"),
			IntraName:  m.pAddForm.GetString("intraName"),
			TeamID:     m.mcontext.CurrentTeamID,
		})
		return switchState(m, TDetailsState)
	}

	return m, tea.Batch(cmds...)
}
