package db

import "core-cli/model"

func SavePlayer(player *model.Player) {
	db.Save(player)
}

func SaveTeam(team *model.Team) {
	db.Save(team)
}

func ToggleTeamSelection(teamID uint) {
	team := GetTeam(teamID)
	team.Selected = !team.Selected
	db.Save(team)
}

func AddContainer(containerID, gameName string) {
	db.Save(&model.Container{ContainerID: containerID, GameName: gameName})
}

func AddNetwork(networkID, gameName string) {
	db.Save(&model.Network{NetworkID: networkID, GameName: gameName})
}
