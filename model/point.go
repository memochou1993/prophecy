package model

import "gorm.io/gorm"

type Point struct {
	gorm.Model
	Name       string     `gorm:"size:255;not null;uniqueIndex;"`
	HouseID    uint       `gorm:"not null;" json:"-"`
	Properties []Property `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"`
	Entries    []Entry    `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"`
}
