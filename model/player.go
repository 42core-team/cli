package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	IntraName  string
	GithubName string
	TeamID     uint
}
