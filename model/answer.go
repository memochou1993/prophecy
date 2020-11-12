package model

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	UserID   uint `gorm:"not null;"`
	ChoiceID uint `gorm:"not null;"`
}
