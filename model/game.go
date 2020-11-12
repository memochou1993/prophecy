package model

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Name         string        `gorm:"size:255;not null;"`
	HouseID      uint          `gorm:"not null;"`
	Questions    []Question    `gorm:"constraint:OnDelete:CASCADE;"`
	Users        []User        `gorm:"many2many:tickets;constraint:OnDelete:CASCADE;"`
	Participants []Participant `gorm:"polymorphic:Owner;polymorphicValue:game;constraint:OnDelete:CASCADE;"`
}
