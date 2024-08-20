package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	IntraName  string
	GithubName string
	GithubID   int64
	TeamID     uint
}
