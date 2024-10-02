package game

import (
	"core-cli/db"
	"core-cli/docker"
	"core-cli/model"
	"os"
)

func RunGame(team1, team2 model.Team) error {
	docker.PullImage(os.Getenv("SERVER_IMAGE"))
	docker.PullImage(os.Getenv("BOT_CLIENT_IMAGE"))

	name := "game-" + team1.Name + "-" + team2.Name

	networkID := docker.CreateNetwork(name)
	db.AddNetwork(networkID, name)

	resp, err := docker.CreateServerContainer("server-"+name, os.Getenv("SERVER_IMAGE"), networkID, []string{
		"./game", "1", "2",
	})
	if err != nil {
		return err
	}
	serverID := resp.ID
	db.AddContainer(serverID, name)

	resp, err = docker.CreateBotContainer("bot1-"+name, os.Getenv("BOT_CLIENT_IMAGE"), networkID, []string{
		"SERVER_IP=server-" + name,
		"REPO_URL=https://github.com/" + os.Getenv("GITHUB_ORG") + "/" + team1.RepoName,
		"GIT_ACCESS_TOKEN=" + os.Getenv("GITHUB_TOKEN"),
		"PLAYER_ID=1",
	})
	if err != nil {
		return err
	}
	bot1ID := resp.ID
	db.AddContainer(bot1ID, name)

	resp, err = docker.CreateBotContainer("bot2-"+name, os.Getenv("BOT_CLIENT_IMAGE"), networkID, []string{
		"SERVER_IP=server-" + name,
		"REPO_URL=https://github.com/" + os.Getenv("GITHUB_ORG") + "/" + team2.RepoName,
		"GIT_ACCESS_TOKEN=" + os.Getenv("GITHUB_TOKEN"),
		"PLAYER_ID=2",
	})
	if err != nil {
		return err
	}
	bot2ID := resp.ID
	db.AddContainer(bot2ID, name)

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

	return nil
}
