package tui

import (
	"core-cli/db"
	"core-cli/game"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

func runSelectedGame() int {
	teams := db.GetSelectedTeams()

	if len(teams) == 0 {
		huh.NewForm(
			huh.NewGroup(
				huh.NewNote().Title("No teams selected").Description("Please go to the team list and select two teams"),
			),
		).Run()
		return Nothing
	} else if len(teams) != 2 {
		huh.NewForm(
			huh.NewGroup(
				huh.NewNote().Title("Select two teams").Description("Please go to the team list and select two teams"),
			),
		).Run()
		return Nothing
	}

	spinner.New().Title("Running game...").Action(func() {
		err := game.RunGame(teams[0], teams[1])
		if err != nil {
			log.Default().Println(err)
		}
	}).Run()

	return Nothing
}
