package model

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Team1ID    uint
	Team1Name  string
	Team2ID    uint
	Team2Name  string
	Winner     uint
	WinnerName string
	Status     string
}
