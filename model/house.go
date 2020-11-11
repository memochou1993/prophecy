package model

import "gorm.io/gorm"

type House struct {
	gorm.Model
	Name     string  `gorm:"size:255;not null;uniqueIndex;"`
	Currency string  `gorm:"size:20;not null;uniqueIndex;"`
	Settings string  `gorm:"type:json;"`
	UserID   uint    `gorm:"not null;"`
	Entries  []Entry `gorm:"constraint:OnDelete:CASCADE;"`
	Games    []Game  `gorm:"constraint:OnDelete:CASCADE;"`
}
