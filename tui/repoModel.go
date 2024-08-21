package tui

import (
	"core-cli/db"
	"core-cli/github"
	"errors"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
)

func runCreateRepos() int {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Create Repos").Description("Enter the url of the template repo").Key("url").Validate(func(input string) error {
				if input == "" {
					return errors.New("url cannot be empty")
				}
				_, err := github.GetRepoFromURL(input)
				if err != nil {
					return err
				}
				return nil
			}),
			huh.NewConfirm().Title("Confirm").Description("Do you want to create repos?").Key("confirm"),
		),
	)

	err := form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return UserAborted
		}
		log.Fatal(err)
	}

	if !form.GetBool("confirm") {
		return GoBack
	}

	templateRepo, err := github.GetRepoFromURL(form.GetString("url"))
	if err != nil {
		log.Fatal(err)
	}

	teams := db.GetTeams()

	for ind, team := range teams {
		ShowLoadingScreen("Creating repos ("+strconv.Itoa(ind+1)+"/"+strconv.Itoa(len(teams))+")", func() {
			_, _ = github.CreateRepoFromTemplate(team.Name, templateRepo)
		})
	}

	return Nothing
}
