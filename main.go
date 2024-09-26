package main

import (
	"core-cli/db"
	"core-cli/docker"
	"core-cli/github"
	"core-cli/logging"
	"core-cli/tui"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	tui.ShowLoadingScreen("Loading...", func() {
		godotenv.Load()
		logging.SetupLogToFile()
	})
	defer logging.CloseLogToFile()

	tui.ShowLoadingScreen("Init Github client...", func() {
		err := github.NewClient()
		if err != nil {
			log.Default().Fatalln("Error creating GitHub client:", err)
		}
	})

	tui.ShowLoadingScreen("Connect to database...", func() {
		db.Connect()
	})

	tui.ShowLoadingScreen("Init Docker client...", func() {
		err := docker.NewDockerClient()
		if err != nil {
			log.Default().Fatalln("Error creating Docker client:", err)
		}
	})
	defer docker.CloseDockerClient()

	tui.Start()
}
