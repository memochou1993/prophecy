package model

import "gorm.io/gorm"

type Participant struct {
	gorm.Model
	UserID    uint `gorm:"not null;"`
	OwnerID   uint
	OwnerType string `gorm:"size:20;not null;"`
}
