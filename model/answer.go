package model

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	UserID   uint `gorm:"not null;" json:"-"`
	ChoiceID uint `gorm:"not null;" json:"-"`
}
