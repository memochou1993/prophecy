package model

import "gorm.io/gorm"

type House struct {
	gorm.Model
	Name         string        `gorm:"size:255;not null;uniqueIndex;"`
	Settings     string        `gorm:"type:json;"`
	UserID       uint          `gorm:"not null;"`
	Point        Point         `gorm:"constraint:OnDelete:CASCADE;"`
	Property     Property      `gorm:"polymorphic:Owner;polymorphicValue:house;"`
	Entries      []Entry       `gorm:"polymorphic:Owner;polymorphicValue:house;"`
	Participants []Participant `gorm:"polymorphic:Owner;polymorphicValue:house;"`
}
