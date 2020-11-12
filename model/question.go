package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Question struct {
	gorm.Model
	ReviewedAt   sql.NullTime
	PublishedAt  sql.NullTime
	OpenedAt     time.Time
	ClosedAt     time.Time
	Title        string `gorm:"size:255;not null;"`
	Description  string
	HouseID      uint          `gorm:"not null;" json:"-"`
	Choices      []Choice      `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"`
	Participants []Participant `gorm:"polymorphic:Owner;polymorphicValue:question;" json:",omitempty"`
	Property     []Property    `gorm:"polymorphic:Owner;polymorphicValue:question;" json:",omitempty"`
	Entries      []Entry       `gorm:"polymorphic:Owner;polymorphicValue:question;" json:",omitempty"`
}
