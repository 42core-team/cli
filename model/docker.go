package model

import (
	"gorm.io/gorm"
)

type Container struct {
	gorm.Model
	ContainerID string `gorm:"uniqueIndex;not null"`
	GameName    string
}

type Network struct {
	gorm.Model
	NetworkID string `gorm:"uniqueIndex;not null"`
	GameName  string
}
