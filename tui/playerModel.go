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
		githubUser, err := github.GetGithubUserByUserName(form.GetString("githubName"))
		if err != nil {
			log.Fatal(err)
		}

		db.SavePlayer(&model.Player{
			GithubName: form.GetString("githubName"),
			IntraName:  form.GetString("intraName"),
			GithubID:   *githubUser.ID,
			TeamID:     uint(teamID),
		})
	})

	return Nothing
}

func runPDetailsForm(playerID int) int {
	player := db.GetPlayer(uint(playerID))
	playerCopy := *player

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
					if input != playerCopy.GithubName && db.PlayerExistsByGithubName(input) {
						return errors.New(input + " already exists in the db")
					}
					if input != playerCopy.GithubName && !github.GithubUserExists(input) {
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
			huh.NewConfirm().Key("delete").Title("Delete Player").Description("Do you want to delete the player?"),
		),
	)

	err := form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return UserAborted
		}
		log.Fatal(err)
	}

	if form.GetBool("save") {
		ShowLoadingScreen("Saving player", func() {
			githubUser, err := github.GetGithubUserByUserName(form.GetString("githubName"))
			if err != nil {
				log.Fatal(err)
			}
			player.GithubID = *githubUser.ID

			db.SavePlayer(player)
		})
	} else if form.GetBool("delete") {
		ShowLoadingScreen("Deleting player", func() {
			db.DeletePlayer(player)
		})
	}

	return Nothing
}
