package tui

import (
	"core-cli/db"
	"core-cli/github"
	"errors"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type State int

const (
	TListState State = iota
	TDetailsState
	PListState
	PAddState
)

type Model struct {
	state        State
	tListForm    *huh.Form
	tDetailsForm *huh.Form
	pListForm    *huh.Form
	pAddForm     *huh.Form
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
				Key("playerList").
				Options(huh.NewOptions("New", "Player 1", "Player 2")...).
				Title("Player List").
				Description("Choose a player to view details"),
		),
	)
	return m.pListForm.Init()
}

func initPAddForm(m *Model) tea.Cmd {
	m.pAddForm = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("githubUserName").
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
				Key("slackUserName").
				Description("Enter the slack username of the player"),
		),
	)
	return m.pAddForm.Init()
}

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
