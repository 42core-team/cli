package db

import "core-cli/model"

func DeletePlayer(player *model.Player) {
	db.Delete(player)
}

func DeleteTeamAndPlayer(team *model.Team) {
	db.Delete(team)
}