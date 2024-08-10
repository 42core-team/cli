package db

import "core-cli/model"

func GetPlayer(id uint) *model.Player {
	var player model.Player
	db.First(&player, id)
	return &player
}

func GetPlayerByIntraName(intraName string) *model.Player {
	var player model.Player
	db.Where("intra_name = ?", intraName).First(&player)
	return &player
}

func GetPlayerByGithubName(githubName string) *model.Player {
	var player model.Player
	db.Where("github_name = ?", githubName).First(&player)
	return &player
}

func GetTeamByName(name string) *model.Team {
	var team model.Team
	db.Where("name = ?", name).First(&team)
	return &team
}

func GetTeamByRepoName(repoName string) *model.Team {
	var team model.Team
	db.Where("repo_name = ?", repoName).First(&team)
	return &team
}

func GetPlayersByTeamID(teamID uint) []model.Player {
	var players []model.Player
	db.Where("team_id = ?", teamID).Find(&players)
	return players
}

func GetPLayersByTeamName(name string) []model.Player {
	team := GetTeamByName(name)
	if team == nil {
		return []model.Player{}
	}
	return GetPlayersByTeamID(team.ID)
}

func GetTeams() []model.Team {
	var teams []model.Team
	db.Find(&teams)
	return teams
}
