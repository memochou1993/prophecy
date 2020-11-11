package model

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	Amount  int64 `gorm:"not null"`
	HouseID uint  `gorm:"not null"`
	UserID  uint  `gorm:"not null"`
}
