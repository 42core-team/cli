package tui

import (
	"core-cli/db"
	"core-cli/model"
	"errors"
	"log"

	"github.com/charmbracelet/huh"
)

func runTList() int {
	var teamID int = 0

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Value(&teamID).
				OptionsFunc(func() []huh.Option[int] {
					var options []huh.Option[int]
					options = append(options, huh.NewOption[int]("<New>", NewEntry))

					for _, team := range db.GetTeams() {
						options = append(options, huh.NewOption(team.Name, int(team.ID)))
					}
					return options
				}, "static").
				Title("Team List").
				Description("Choose a team to view details or create a new one"),
		),
	)

	err := form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return UserAborted
		}
		log.Fatal(err)
	}

	return teamID
}

func runTAddForm() int {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("name").
				Title("Add Team").
				Description("Enter the name of the team").
				Validate(func(input string) error {
					if input == "" {
						return errors.New("team name cannot be empty")
					}
					if db.TeamExistsByName(input) {
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

	ShowLoadingScreen("Adding team", func() {
		db.SaveTeam(&model.Team{
			Name:     form.GetString("name"),
			RepoName: form.GetString("repoName"),
		})
	})

	return Nothing
}

func runTDetails(teamID int) int {
	var playerID int = GoBack

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Value(&playerID).
				Key("teamDetails").
				Title("Team Details").
				Description("Choose an option").
				OptionsFunc(func() []huh.Option[int] {
					var options []huh.Option[int]
					options = append(options, huh.NewOption[int]("<Back>", GoBack))
					options = append(options, huh.NewOption[int]("<New>", NewEntry))
					options = append(options, huh.NewOption[int]("<Delete>", DeleteEntry))

					for _, player := range db.GetPlayersByTeamID(uint(teamID)) {
						options = append(options, huh.NewOption(player.IntraName, int(player.ID)))
					}
					return options
				}, &teamID),
		),
	)

	err := form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return UserAborted
		}
		log.Fatal(err)
	}

	return playerID
}

func runTDelete(teamID int) int {
	team := db.GetTeam(uint(teamID))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Delete Team " + team.Name).
				Description("Do you really want to delete the team and all of its players?").
				Key("delete"),
		),
	)

	err := form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return UserAborted
		}
		log.Fatal(err)
	}

	if form.GetBool("delete") {
		ShowLoadingScreen("Deleting team", func() {
			db.DeleteTeamAndPlayer(team)
		})
		return Success
	}

	return Nothing
}
