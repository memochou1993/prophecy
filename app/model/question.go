package model

import (
	"gorm.io/gorm"
	"time"
)

type Question struct {
	gorm.Model
	ReviewedAt   NullTime
	PublishedAt  NullTime
	OpenedAt     time.Time
	ClosedAt     time.Time
	Title        string `gorm:"size:255;not null;"`
	Description  string
	HouseID      uint          `gorm:"not null;" json:"-"`
	Choices      []Choice      `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"`
	Users        []User        `gorm:"many2many:question;" json:",omitempty"`
	Participants []Participant `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"` // constraint
	Property     *Property     `gorm:"polymorphic:Owner;polymorphicValue:question;" json:",omitempty"`
	Entries      []Entry       `gorm:"polymorphic:Owner;polymorphicValue:question;" json:",omitempty"`
}
