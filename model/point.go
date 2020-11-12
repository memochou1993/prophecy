package model

import "gorm.io/gorm"

type Point struct {
	gorm.Model
	Name       string     `gorm:"size:255;not null;uniqueIndex;"`
	HouseID    uint       `gorm:"not null;"`
	Properties []Property `gorm:"constraint:OnDelete:CASCADE;"`
	Entries    []Entry    `gorm:"constraint:OnDelete:CASCADE;"`
}
