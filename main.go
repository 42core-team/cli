package main

import (
	"core-cli/db"
	"core-cli/github"
	"core-cli/logging"
	"core-cli/tui"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logging.SetupLogToFile()
	defer logging.CloseLogToFile()

	err := github.NewClient()
	if err != nil {
		log.Fatalln("Error creating GitHub client:", err)
	}
	db.Connect()

	_, err = tea.NewProgram(tui.NewModel()).Run()
	if err != nil {
		log.Fatalln("Error running program:", err)
	}
}
