package model

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Amount    int64 `gorm:"not null;"`
	HouseID   uint  `gorm:"not null;"`
	OwnerID   uint
	OwnerType string `gorm:"size:20;not null;"`
}
