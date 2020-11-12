package model

import (
	"gorm.io/gorm"
	"time"
)

type Question struct {
	gorm.Model
	ReviewedAt   time.Time
	PublishedAt  time.Time
	OpenedAt     time.Time
	ClosedAt     time.Time
	Title        string `gorm:"size:255;not null;"`
	Description  string
	Choices      []Choice      `gorm:"constraint:OnDelete:CASCADE;"`
	Property     []Property    `gorm:"polymorphic:Owner;polymorphicValue:question;"`
	Entries      []Entry       `gorm:"polymorphic:Owner;polymorphicValue:question;"`
	Participants []Participant `gorm:"polymorphic:Owner;polymorphicValue:question;"`
}
