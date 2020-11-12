package model

import "gorm.io/gorm"

type Prize struct {
	gorm.Model
	Name      string `gorm:"size:255;not null;"`
	Cost      int64  `gorm:"not null;"`
	OwnerID   uint
	OwnerType string `gorm:"size:20;not null;"`
}
