package tui

func Start() {
	for {
		teamID := runTList()
		if teamID == -1 {
			break
		}
		runTDetails(uint(teamID))
	}
}
