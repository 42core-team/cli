package main

import (
	"core-cli/github"
	"core-cli/tui"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	github.NewClient()
	_, err := tea.NewProgram(tui.NewModel()).Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
}
