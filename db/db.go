package db

import (
	"core-cli/github"
	"core-cli/model"
	"log"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Connect() {
	if os.Getenv("GITHUB_USE_DB_REPO") == "true" {
		pullDatabase()
	}

	var err error
	db, err = gorm.Open(sqlite.Open("./cli-db/coreEvent.db"), &gorm.Config{
		Logger: logger.New(
			log.Default(),
			logger.Config{
				SlowThreshold: 0,
				LogLevel:      logger.Info,
				Colorful:      false,
			},
		),
	})
	if err != nil {
		log.Default().Fatal(err)
	}

	db.AutoMigrate(model.Player{}, model.Team{})

	if os.Getenv("GITHUB_USE_DB_REPO") == "true" {
		pushDatabase()
	}
}

func pullDatabase() {
	err := github.Pull("./cli-db")
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "repository does not exist") {
			repo, err := github.GetRepoFromName("cli-db")
			if err != nil {
				repo, err = github.CreateRepo("cli-db")
				if err != nil {
					log.Default().Fatal(err)
				}
			}

			_, err = github.Clone(*repo.CloneURL, "./cli-db")
			if err != nil {
				if strings.Contains(err.Error(), "remote repository is empty") {
					os.Mkdir("./cli-db", 0755)
				} else {
					log.Default().Fatal(err)
				}
			}
		} else if !strings.Contains(err.Error(), "remote repository is empty") {
			log.Default().Fatal(err)
		}
	}
}

func pushDatabase() {
	err := github.CommitAndPush("./cli-db", "Update database")
	if err != nil {
		log.Default().Fatal(err)
	}
}
