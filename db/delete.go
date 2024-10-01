package db

import "core-cli/model"

func DeletePlayer(player *model.Player) {
	db.Delete(player)
}

func DeleteTeamAndPlayer(team *model.Team) {
	db.Delete(team)
	db.Delete(&model.Player{}, "team_id = ?", team.ID)
}

func DeleteContainer(containerID string) {
	db.Delete(&model.Container{}, "container_id = ?", containerID)
}

func DeleteNetwork(networkID string) {
	db.Delete(&model.Network{}, "network_id = ?", networkID)
}

func DeleteAllContainers() {
	db.Delete(&model.Container{})
}

func DeleteAllNetworks() {
	db.Delete(&model.Network{})
}
