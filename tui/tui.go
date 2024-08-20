package tui

const (
	UserAborted = -1
	NewEntry    = 0
)

func Start() {
Loop:
	for {
		teamID := runTList()
		switch teamID {
		case UserAborted:
			break Loop
		case NewEntry:
			// runTAddForm()
		default:
			runTDetails(uint(teamID))
		}
	}
}
