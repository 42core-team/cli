package tui

import (
	"errors"
	"log"

	"github.com/charmbracelet/huh"
)

func runMain() int {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Options(
				huh.NewOption("Team List", "teamlist"),
				huh.NewOption("Create Repos", "createrepos"),
				huh.NewOption("Rm write Access", "rmwriteaccess"),
				huh.NewOption("Run selected game", "runselectedgame"),
				huh.NewOption("Cleanup Docker", "cleanupdocker"),
			).Title("Main Menu").Description("Choose an option").Key("main"),
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
	case "rmwriteaccess":
		runRemoveWriteAccess()
		return runMain()
	case "runselectedgame":
		runSelectedGame()
		return runMain()
	case "cleanupdocker":
		runCleanupDocker()
		return runMain()
	default:
		return Nothing
	}
}