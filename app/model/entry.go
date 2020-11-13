package model

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	Amount    int64  `gorm:"not null;"`
	PointID   uint   `gorm:"not null;" json:"-"`
	OwnerID   uint   `json:"-"`
	OwnerType string `gorm:"size:20;not null;" json:"-"`
	Point     Point
}
