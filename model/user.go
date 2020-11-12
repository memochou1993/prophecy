package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `gorm:"size:255;not null;"`
	Email    string  `gorm:"size:255;not null;uniqueIndex;"`
	Password string  `gorm:"size:255;not null;"`
	House    House   `gorm:"constraint:OnDelete:CASCADE;"`
	Entries  []Entry `gorm:"polymorphic:Owner;polymorphicValue:user;constraint:OnDelete:CASCADE;"`
}
