package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `gorm:"size:255;not null"`
	Email    string  `gorm:"size:255;not null;uniqueIndex"`
	Password string  `gorm:"size:255;not null"`
	Houses   []House `gorm:"constraint:OnDelete:CASCADE;foreignKey:OwnerID"`
	Entries  []Entry `gorm:"constraint:OnDelete:CASCADE"`
}
