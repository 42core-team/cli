package db

import "core-cli/model"

func SavePlayer(player *model.Player) {
	db.Save(player)
}

func SaveTeam(team *model.Team) {
	db.Save(team)
}
