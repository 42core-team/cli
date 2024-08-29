package tui

import (
	"errors"
	"log"

	"github.com/charmbracelet/huh"
)

func runMain() int {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Options(huh.NewOption("Team List", "teamlist"), huh.NewOption("Create Repos", "createrepos")).Title("Main Menu").Description("Choose an option").Key("main"),
		),
	)

	err := form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return UserAborted
		}
		log.Default().Fatal(err)
	}

	switch form.GetString("main") {
	case "teamlist":
		handleTList()
		return Nothing
	case "createrepos":
		runCreateRepos()
		return runMain()
	default:
		return Nothing
	}
}
