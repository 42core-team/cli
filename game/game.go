package game

import (
	"core-cli/docker"
	"core-cli/model"
	"os"
)

func RunGame(team1, team2 model.Team) error {
	err := docker.PullImage(os.Getenv("SERVER_IMAGE"))
	if err != nil {
		return err
	}
	err = docker.PullImage(os.Getenv("BOT_CLIENT_IMAGE"))
	if err != nil {
		return err
	}

	name := "game-" + team1.Name + "-" + team2.Name

	networkID := docker.CreateNetwork(name)

	resp, err := docker.CreateServerContainer("server-"+name, os.Getenv("SERVER_IMAGE"), networkID, []string{
		"./game", "1", "2",
	})
	if err != nil {
		return err
	}
	serverID := resp.ID

	resp, err = docker.CreateBotContainer("bot1-"+name, os.Getenv("BOT_CLIENT_IMAGE"), networkID, []string{
		"SERVER_IP=server-" + name,
		"REPO_URL=" + team1.RepoName,
		"GIT_ACCESS_TOKEN=" + os.Getenv("GITHUB_TOKEN"),
		"PLAYER_ID=1",
	})
	if err != nil {
		return err
	}
	bot1ID := resp.ID

	resp, err = docker.CreateBotContainer("bot2-"+name, os.Getenv("BOT_CLIENT_IMAGE"), networkID, []string{
		"SERVER_IP=server-" + name,
		"REPO_URL=" + team2.RepoName,
		"GIT_ACCESS_TOKEN=" + os.Getenv("GITHUB_TOKEN"),
		"PLAYER_ID=2",
	})
	if err != nil {
		return err
	}
	bot2ID := resp.ID

	err = docker.StartContainer(serverID)
	if err != nil {
		return err
	}
	err = docker.StartContainer(bot1ID)
	if err != nil {
		return err
	}
	err = docker.StartContainer(bot2ID)
	if err != nil {
		return err
	}
}
