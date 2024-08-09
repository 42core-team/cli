package db

import (
	"core-cli/github"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	pullDatabase()
	var err error
	db, err = gorm.Open(sqlite.Open("./cli-db/coreEvent.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func pullDatabase() {
	err := github.Pull("./cli-db")
	if err != nil {
		repo, err := github.CreateRepo("cli-db")
		if err != nil {

			panic(err)
		}

		_, err = github.Clone(*repo.CloneURL, "./cli-db")
		if err != nil {
			panic(err)
		}
	}
}

func pushDatabase() {
	err := github.CommitAndPush("cli-db", "Update database")
	if err != nil {
		panic(err)
	}
}
