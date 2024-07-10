package main

import (
	"core-cli/github"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	err := github.NewClient()
	if err != nil {
		log.Fatalln("Error creating GitHub client:", err)
	}
	github.CheckGithubUsername("paulicen")
	// _, err = tea.NewProgram(tui.NewModel()).Run()
	// if err != nil {
	// 	log.Fatalln("Error running program:", err)
	// }
}
