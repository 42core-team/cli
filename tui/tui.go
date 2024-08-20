package tui

const (
	Nothing     = -3
	UserAborted = -2
	GoBack      = -1
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
			handleTDetails(teamID)
		}
	}
}

func handleTDetails(teamID int) {
Loop:
	for {
		playerID := runTDetails(teamID)
		switch playerID {
		case UserAborted:
			break Loop
		case GoBack:
			break Loop
		case NewEntry:
			runPAddForm(teamID)
		default:
			runPDetailsForm(playerID)
		}
	}
}
