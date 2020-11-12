package model

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	Amount    int64 `gorm:"not null;"`
	PointID   uint  `gorm:"not null;"`
	OwnerID   uint
	OwnerType string `gorm:"size:20;not null;"`
}
