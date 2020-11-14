package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	UserID  uint `gorm:"not null;" json:"-"`
	HouseID uint `gorm:"not null;" json:"-"`
	IsOwner bool `gorm:"not null;"`
}
