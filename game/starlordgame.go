package game

import (
	"core-cli/db"
	"core-cli/docker"
	"core-cli/github"
	"core-cli/model"
	"log"
	"os"
	"time"
)

func RunGameAgainstStarlord(team *model.Team) error {
	// docker.PullImage(os.Getenv("SERVER_IMAGE"))
	// docker.PullImage(os.Getenv("BOT_CLIENT_IMAGE"))

	name := "game-starlord-" + team.Name

	networkID := docker.CreateNetwork(name)
	db.AddNetwork(networkID, name)

	resp, err := docker.CreateServerContainer("server-"+name, os.Getenv("SERVER_IMAGE"), networkID, []string{
		"./game", "1", "2",
	}, []string{
		"TICK_RATE=" + os.Getenv("TICK_RATE"),
	})
	if err != nil {
		return err
	}
	serverID := resp.ID
	db.AddContainer(serverID, name)

	resp, err = docker.CreateBotContainer("bot-"+name, os.Getenv("BOT_CLIENT_IMAGE"), networkID, []string{
		"SERVER_IP=server-" + name,
		"REPO_URL=https://github.com/" + os.Getenv("GITHUB_ORG") + "/" + team.RepoName,
		"GIT_ACCESS_TOKEN=" + os.Getenv("GITHUB_TOKEN"),
		"PLAYER_ID=1",
	})
	if err != nil {
		return err
	}
	bot1ID := resp.ID
	db.AddContainer(bot1ID, name)

	resp, err = docker.CreateBotContainer("starlord-"+name, os.Getenv("BOT_CLIENT_IMAGE"), networkID, []string{
		"SERVER_IP=server-" + name,
		"REPO_URL=" + os.Getenv("STARLORD_REPO"),
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

	go func() {
		for {
			status, err := docker.CheckContainerStatus(bot1ID)
			if err != nil {
				log.Default().Println(err)
				return
			}
			if !status.Running {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}

		logs, err := docker.GetLogs(bot1ID)
		if err != nil {
			log.Default().Println(err)
			return
		}

		_, err = github.CreateTraceRelease(team.RepoName, "```\n"+logs+"\n```", false, false)
		if err != nil {
			log.Default().Println(err)
			return
		}

		docker.StopRmContainer(serverID)
		docker.StopRmContainer(bot1ID)
		docker.StopRmContainer(bot2ID)
		docker.RemoveNetwork(networkID)
	}()

	return nil
}
