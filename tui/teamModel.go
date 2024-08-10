package tui

import (
	"core-cli/db"
	"core-cli/github"
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

func initTDetailsForm(m *Model) tea.Cmd {
	var nameList []string = []string{"<New>"}

	for _, player := range db.GetPlayersByTeamID(0) {
		nameList = append(nameList, player.IntraName)
	}

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
					if !github.GithubUserExists(input) {
						return errors.New(input + " does not exist on github")
					}
					return nil
				}),
			huh.NewInput().
				Key("intraName").
				Description("Enter the intra username of the player"),
		),
	)
	return m.pAddForm.Init()
}
