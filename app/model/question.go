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
	HouseID      uint `gorm:"not null;" json:"-"`
	OwnerID      uint `gorm:"not null;" json:"-"`
	Owner        User
	Users        []User        `gorm:"many2many:participants;" json:",omitempty"`
	Participants []Participant `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"` // constraint
	Choices      []Choice      `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"`
	Property     *Property     `gorm:"polymorphic:Owner;polymorphicValue:question;" json:",omitempty"`
	Entries      []Entry       `gorm:"polymorphic:Owner;polymorphicValue:question;" json:",omitempty"`
}
