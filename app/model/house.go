package model

import "gorm.io/gorm"

type House struct {
	gorm.Model
	Name         string        `gorm:"size:255;not null;uniqueIndex;"`
	Settings     string        `gorm:"type:json;"`
	UserID       uint          `gorm:"not null;" json:"-"`
	Point        *Point        `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"`
	Questions    []Question    `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"`
	Participants []Participant `gorm:"polymorphic:Owner;polymorphicValue:house;" json:",omitempty"`
	Property     *Property     `gorm:"polymorphic:Owner;polymorphicValue:house;" json:",omitempty"`
	Entries      []Entry       `gorm:"polymorphic:Owner;polymorphicValue:house;" json:",omitempty"`
}
