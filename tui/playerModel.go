package tui

import (
	"core-cli/db"
	"core-cli/github"
	"core-cli/model"
	"errors"
	"log"

	"github.com/charmbracelet/huh"
)

func runPAddForm(teamID int) int {
	form := huh.NewForm(
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

	err := form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return UserAborted
		}
		log.Fatal(err)
	}

	ShowLoadingScreen("Adding player", func() {
		db.SavePlayer(&model.Player{
			GithubName: form.GetString("githubName"),
			IntraName:  form.GetString("intraName"),
			TeamID:     uint(teamID),
		})
	})

	return Nothing
}

// func updatePAddForm(m *Model, msg *tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmds []tea.Cmd

// 	form, cmd := m.pAddForm.Update(*msg)
// 	if f, ok := form.(*huh.Form); ok {
// 		m.pAddForm = f
// 		cmds = append(cmds, cmd)
// 	}

// 	if m.pAddForm.State == huh.StateCompleted {
// 		db.SavePlayer(&model.Player{
// 			GithubName: m.pAddForm.GetString("githubName"),
// 			IntraName:  m.pAddForm.GetString("intraName"),
// 			TeamID:     m.mcontext.CurrentTeamID,
// 		})
// 		return switchState(m, TDetailsState)
// 	}

// 	return m, tea.Batch(cmds...)
// }

func runPDetailsForm(playerID uint) error {
	player := db.GetPlayer(playerID)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("githubName").
				Title("Github Name").
				Value(&player.GithubName).
				Validate(func(input string) error {
					if input == "" {
						return errors.New("player name cannot be empty")
					}
					if input != player.GithubName && db.PlayerExistsByGithubName(input) {
						return errors.New(input + " already exists in the db")
					}
					if input != player.GithubName && !github.GithubUserExists(input) {
						return errors.New(input + " does not exist on github")
					}
					return nil
				}),
			huh.NewInput().
				Key("intraName").
				Title("Intra Name").
				Value(&player.IntraName).
				Validate(func(input string) error {
					if input == "" {
						return errors.New("player name cannot be empty")
					}
					if input != player.IntraName && db.PlayerExistsByIntraName(input) {
						return errors.New(input + " already exists in the db")
					}
					return nil
				}),
			huh.NewConfirm().Key("save").Title("Save Changes").Description("Do you want to save the changes?"),
		),
	)

	return form.Run()
}

// func updatePDetailsForm(m *Model, msg *tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmds []tea.Cmd

// 	form, cmd := m.pDetailsForm.Update(*msg)
// 	if f, ok := form.(*huh.Form); ok {
// 		m.pDetailsForm = f
// 		cmds = append(cmds, cmd)
// 	}

// 	if m.pDetailsForm.State == huh.StateCompleted {
// 		if m.pDetailsForm.GetBool("save") {
// 			player := db.GetPlayer(m.mcontext.CurrentPlayerID)
// 			player.GithubName = m.pDetailsForm.GetString("githubName")
// 			player.IntraName = m.pDetailsForm.GetString("intraName")
// 			db.SavePlayer(player)
// 		}
// 		return switchState(m, TDetailsState)
// 	}

// 	return m, tea.Batch(cmds...)
// }
