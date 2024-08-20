package tui

import (
	"core-cli/db"
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
