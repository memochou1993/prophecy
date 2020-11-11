package model

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	UserID uint `gorm:"not null;"`
	GameID uint `gorm:"not null;"`
}
