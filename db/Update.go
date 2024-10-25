package db

func ClearTeamSelections() {
	teams := GetSelectedTeams()
	for _, team := range teams {
		team.Selected = false
		SaveTeam(&team)
	}
}
