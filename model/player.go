package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	IntraName  string
	GithubName string
	TeamID     uint
}

type Team struct {
	gorm.Model
	Name     string
	RepoName string
	Players  []Player
}
