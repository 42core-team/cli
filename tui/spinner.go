package tui

import (
	"log"

	"github.com/charmbracelet/huh/spinner"
)

func ShowLoadingScreen(msg string, action func()) {
	err := spinner.New().Title(msg).Action(action).Run()
	if err != nil {
		log.Fatal(err)
	}
}
