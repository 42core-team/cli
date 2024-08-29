package tui

import (
	"core-cli/db"
	"core-cli/github"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
)

func runCreateRepos() int {
	defaultRepo := os.Getenv("DEFAULT_TEMPLATE_REPO")
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Create Repos").Description("Enter the url of the template repo").Key("url").Value(&defaultRepo).Validate(func(input string) error {
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
		log.Default().Fatal(err)
	}

	if !form.GetBool("confirm") {
		return GoBack
	}

	templateRepo, err := github.GetRepoFromURL(form.GetString("url"))
	if err != nil {
		log.Default().Fatal(err)
	}

	teams := db.GetTeams()

	for ind, team := range teams {
		ShowLoadingScreen("Creating repos ("+strconv.Itoa(ind+1)+"/"+strconv.Itoa(len(teams))+")", func() {
			repo, err := github.GetRepoFromName(team.Name)
			if err != nil {
				repo, err = github.CreateRepoFromTemplate(team.Name, templateRepo)
				if err != nil {
					log.Default().Println(err)
					return
				}
				team.RepoName = *repo.Name
				db.SaveTeam(&team)
			}

			for _, player := range team.Players {
				err = github.AddCollaborator(*repo.Name, player.GithubName)
				if err != nil {
					log.Default().Println(err)
					return
				}
			}
		})
	}

	return Nothing
}

func runRemoveWriteAccess() int {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Remove Write Access").Description("Do you want to remove write access from all repos?").Key("confirm"),
		),
	)

	err := form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return UserAborted
		}
		log.Default().Fatal(err)
	}

	if !form.GetBool("confirm") {
		return GoBack
	}

	teams := db.GetTeams()

	for ind, team := range teams {
		ShowLoadingScreen("Removing write access ("+strconv.Itoa(ind+1)+"/"+strconv.Itoa(len(teams))+")", func() {
			github.ChangeCollaboratorInviteReadOnly(team.RepoName, team.RepoName)

			for _, player := range team.Players {
				err = github.ChangeCollaboratorReadOnly(team.RepoName, player.GithubName)
				if err != nil {
					log.Default().Println(err)
					return
				}
			}
		})
	}

	return Nothing
}
