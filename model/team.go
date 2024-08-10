package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name     string
	RepoName string
	Players  []Player
}
