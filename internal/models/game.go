package models

import (
	"gorm.io/gorm"
)

type Game struct {
	ID          uint
	Slug        string  `gorm:"uniqueIndex:index_game_slug"`
	Name        string  `gorm:"not null"`
	Description *string `gorm:"type:text"`
	Image       string  `gorm:"not null"`
	IsActive    bool    `gorm:"not null;default:false"`
	gorm.Model
}
