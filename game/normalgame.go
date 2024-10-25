package game

import (
	"core-cli/db"
	"core-cli/docker"
	"core-cli/model"
	"core-cli/utils"
	"log"
	"os"
	"time"
)

func RunGame(team1, team2 model.Team) error {
	// docker.PullImage(os.Getenv("SERVER_IMAGE"))
	// docker.PullImage(os.Getenv("BOT_CLIENT_IMAGE"))
	// docker.PullImage(os.Getenv("VISUALIZER_IMAGE"))

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

	go func() {
		for {
			status, err := docker.CheckContainerStatus(serverID)
			if err != nil {
				log.Default().Println(err)
				return
			}
			if !status.Running {
				break
			}

			statusBot1, err := docker.CheckContainerStatus(bot1ID)
			if err != nil {
				log.Default().Println(err)
				return
			}
			statusBot2, err := docker.CheckContainerStatus(bot2ID)
			if err != nil {
				log.Default().Println(err)
				return
			}

			if !statusBot1.Running && !statusBot2.Running {
				break
			} else if !statusBot1.Running {
				log.Default().Println("Winner is", team2.Name, " in game ", name, " as ", team1.Name, " bot is not running")
				db.AddGame(&model.Game{
					Team1ID:    team1.ID,
					Team1Name:  team1.Name,
					Team2ID:    team2.ID,
					Team2Name:  team2.Name,
					Winner:     team2.ID,
					WinnerName: team2.Name,
					Status:     "Bot1 not running",
				})
				return
			} else if !statusBot2.Running {
				log.Default().Println("Winner is", team1.Name, " in game ", name, " as ", team2.Name, " bot is not running")
				db.AddGame(&model.Game{
					Team1ID:    team1.ID,
					Team1Name:  team1.Name,
					Team2ID:    team2.ID,
					Team2Name:  team2.Name,
					Winner:     team1.ID,
					WinnerName: team1.Name,
					Status:     "Bot2 not running",
				})
				return
			}

			time.Sleep(500 * time.Millisecond)
		}

		logs, err := docker.GetLogs(serverID)
		if err != nil {
			log.Default().Println(err)
			return
		}

		winner, err := utils.ExtractWinner(logs)
		if err != nil {
			log.Default().Println(err)
			return
		}

		if winner == 1 {
			log.Default().Println("Winner is", team1.Name, " in game ", name)
			db.AddGame(&model.Game{
				Team1ID:    team1.ID,
				Team1Name:  team1.Name,
				Team2ID:    team2.ID,
				Team2Name:  team2.Name,
				Winner:     team1.ID,
				WinnerName: team1.Name,
				Status:     "Success",
			})
		} else if winner == 2 {
			log.Default().Println("Winner is", team2.Name, " in game ", name)
			db.AddGame(&model.Game{
				Team1ID:    team1.ID,
				Team1Name:  team1.Name,
				Team2ID:    team2.ID,
				Team2Name:  team2.Name,
				Winner:     team2.ID,
				WinnerName: team2.Name,
				Status:     "Success",
			})
		}

		docker.StopRmContainer(serverID)
		docker.StopRmContainer(bot1ID)
		docker.StopRmContainer(bot2ID)
		docker.RemoveNetwork(networkID)
	}()

	return nil
}
