package tui

import "core-cli/github"

func Test() {
	github.CreateRepo("test")
	github.AddCollaborator("test", "61714149")
}
