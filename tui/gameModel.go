package tui

import (
	"core-cli/db"
	"core-cli/docker"
	"core-cli/game"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

func runSelectedGame() int {
	teams := db.GetSelectedTeams()

	msgTitle := ""
	msgDesc := ""

	if len(teams) == 0 {
		msgTitle = "No teams selected"
		msgDesc = "Please go to the team list and select two teams"
	} else if len(teams) != 2 {
		msgTitle = "Select two teams"
		msgDesc = "Please go to the team list and select two teams"
	} else if teams[0].RepoName == "" || teams[1].RepoName == "" {
		msgTitle = "No repo found"
		msgDesc = "Please go to the team list and create repos for the selected teams"
	}

	if msgTitle != "" {
		huh.NewForm(
			huh.NewGroup(
				huh.NewNote().Title(msgTitle).Description(msgDesc),
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

func runAgainstStarlord(teamID uint) int {
	spinner.New().Title("Running game...").Action(func() {
		err := game.RunGameAgainstStarlord(db.GetTeam(teamID))
		if err != nil {
			log.Default().Println(err)
		}
	}).Run()

	return Nothing
}

func runCleanupDocker() int {
	spinner.New().Title("Cleaning up...").Action(func() {
		docker.CleanUP()
	}).Run()

	return Nothing
}

func runTraces() int {
	spinner.New().Title("Starting games...").Action(func() {
		for _, team := range db.GetTeams() {
			err := game.RunGameAgainstStarlord(&team)
			if err != nil {
				log.Default().Println(err)
			}
		}
	}).Run()

	return Nothing
}
