package db

import (
	"core-cli/github"
	"os"
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
	pushDatabase()
}

func pullDatabase() {
	err := github.Pull("./cli-db")
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "repository does not exist") {
			repo, err := github.GetRepo("cli-db")
			if err != nil {
				repo, err = github.CreateRepo("cli-db")
				if err != nil {
					panic(err)
				}
			}

			_, err = github.Clone(*repo.CloneURL, "./cli-db")
			if err != nil {
				if strings.Contains(err.Error(), "remote repository is empty") {
					os.Mkdir("./cli-db", 0755)
				} else {
					panic(err)
				}
			}
		} else if !strings.Contains(err.Error(), "remote repository is empty") {
			panic(err)
		}
	}
}

func pushDatabase() {
	err := github.CommitAndPush("./cli-db", "Update database")
	if err != nil {
		panic(err)
	}
}
