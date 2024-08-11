package tui

import "log"

func Start() {
	err := runTListModel()
	if err != nil {
		log.Fatal(err)
	}
}
