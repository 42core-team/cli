package tui

import "core-cli/db"

const (
	Clear              = -9
	RunAgainstStarlord = -8
	Nothing            = -7
	Select             = -6
	Reset              = -5
	UserAborted        = -4
	GoBack             = -3
	Success            = -2
	DeleteEntry        = -1
	NewEntry           = 0
)

func Start() {
	for {
		if runMain() == UserAborted {
			break
		}
	}
}

func handleTList() {
Loop:
	for {
		teamID := runTList()
		switch teamID {
		case UserAborted:
			break Loop
		case GoBack:
			break Loop
		case NewEntry:
			runTAddForm()
		case Clear:
			runTClearSelection()
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
		case DeleteEntry:
			switch runTDelete(teamID) {
			case Success:
				break Loop
			}
		case Reset:
			runTRepoReset(teamID)
		case Select:
			db.ToggleTeamSelection(uint(teamID))
			continue Loop
		case RunAgainstStarlord:
			runAgainstStarlord(uint(teamID))
		default:
			runPDetailsForm(playerID)
		}
	}
}
