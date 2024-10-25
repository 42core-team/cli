package tui

import (
	"core-cli/db"
	"core-cli/github"
	"core-cli/model"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
)

func runTList() int {
	var teamID int = 0

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Value(&teamID).
				OptionsFunc(func() []huh.Option[int] {
					var options []huh.Option[int]
					options = append(options, huh.NewOption[int]("<Back>", GoBack))
					options = append(options, huh.NewOption[int]("<New>", NewEntry))
					options = append(options, huh.NewOption[int]("<Clear Selection>", Clear))

					for _, team := range db.GetTeams() {
						options = append(options, huh.NewOption(team.Name, int(team.ID)))
					}
					return options
				}, "static").
				Title("Team List").
				Description("Choose a team to view details or create a new one"),
		),
	)

	err := form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return UserAborted
		}
		log.Default().Fatal(err)
	}

	return teamID
}

func runTAddForm() int {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("name").
				Title("Add Team").
				Description("Enter the name of the team").
				Validate(func(input string) error {
					if input == "" {
						return errors.New("team name cannot be empty")
					}
					if strings.ContainsAny(input, " ") {
						return errors.New("team name cannot contain spaces")
					}
					if db.TeamExistsByName(input) {
						return errors.New(input + " already exists in the db")
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return UserAborted
		}
		log.Default().Fatal(err)
	}

	ShowLoadingScreen("Adding team", func() {
		db.SaveTeam(&model.Team{
			Name:     form.GetString("name"),
			RepoName: form.GetString("repoName"),
		})
	})

	return Nothing
}

func runTClearSelection() int {
	ShowLoadingScreen("Clearing selection", func() {
		db.ClearTeamSelections()
	})

	return Nothing
}

func runTDetails(teamID int) int {
	var playerID int = GoBack

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Value(&playerID).
				Key("teamDetails").
				Title("Team Details").
				Description("Choose an option").
				OptionsFunc(func() []huh.Option[int] {
					team := db.GetTeam(uint(teamID))
					var options []huh.Option[int]

					options = append(options, huh.NewOption[int]("<Back>", GoBack))
					options = append(options, huh.NewOption[int]("<New>", NewEntry))
					options = append(options, huh.NewOption[int]("<Delete>", DeleteEntry))

					if team.Selected {
						options = append(options, huh.NewOption[int]("<[X] Deselect>", Select))
					} else {
						options = append(options, huh.NewOption[int]("<[ ] Select>", Select))
					}

					options = append(options, huh.NewOption[int]("<Reset Repo>", Reset))
					options = append(options, huh.NewOption[int]("<Run against Starlord>", RunAgainstStarlord))

					for _, player := range db.GetPlayersByTeamID(uint(teamID)) {
						options = append(options, huh.NewOption(player.IntraName+" - "+player.GithubName, int(player.ID)))
					}
					return options
				}, &teamID),
		),
	)

	err := form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return UserAborted
		}
		log.Default().Fatal(err)
	}

	return playerID
}

func runTDelete(teamID int) int {
	team := db.GetTeam(uint(teamID))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Delete Team " + team.Name).
				Description("Do you really want to delete the team and all of its players?").
				Key("delete"),
		),
	)

	err := form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return UserAborted
		}
		log.Default().Fatal(err)
	}

	if form.GetBool("delete") {
		ShowLoadingScreen("Deleting team", func() {
			db.DeleteTeamAndPlayer(team)
		})
		return Success
	}

	return Nothing
}

func runTRepoReset(teamID int) int {
	team := db.GetTeam(uint(teamID))

	defaultRepo := os.Getenv("DEFAULT_TEMPLATE_REPO")
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Reset repo of team "+team.Name).Description("Enter the url of the template repo").Key("url").Value(&defaultRepo).Validate(func(input string) error {
				if input == "" {
					return errors.New("url cannot be empty")
				}
				_, err := github.GetRepoFromURL(input)
				if err != nil {
					return err
				}
				return nil
			}),
			huh.NewConfirm().Title("Confirm").Description("Do you want to reset the repo?").Key("confirm"),
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

	ShowLoadingScreen("Resetting Repo...", func() {
		github.DeleteRepo(team.RepoName)

		repo, err := github.CreateRepoFromTemplate(team.Name, templateRepo)
		if err != nil {
			log.Default().Println(err)
			return
		}

		team.RepoName = *repo.Name
		db.SaveTeam(team)

		for _, player := range team.Players {
			err = github.AddCollaborator(*repo.Name, player.GithubName)
			if err != nil {
				log.Default().Println(err)
				return
			}
		}
	})

	return Nothing
}
