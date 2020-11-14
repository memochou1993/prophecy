package model

import (
	"time"
)

type Question struct {
	Model
	ReviewedAt   NullTime      `json:"reviewedAt"`
	PublishedAt  NullTime      `json:"publishedAt"`
	OpenedAt     time.Time     `json:"openedAt"`
	ClosedAt     time.Time     `json:"closedAt"`
	Title        string        `gorm:"size:255;not null;" json:"title"`
	Description  string        `gorm:"size:255;not null;" json:"description"`
	HouseID      uint          `gorm:"not null;" json:"houseID" validate:"required"`
	OwnerID      uint          `gorm:"not null;" json:"ownerID"`
	Owner        *User         `gorm:"constraint:OnDelete:CASCADE;" json:"owner,omitempty"`
	Users        []User        `gorm:"many2many:participants;" json:"users,omitempty"`
	Participants []Participant `gorm:"constraint:OnDelete:CASCADE;" json:"-"` // constraint
	Choices      []Choice      `gorm:"constraint:OnDelete:CASCADE;" json:"choices,omitempty"`
	Property     *Property     `gorm:"polymorphic:Owner;polymorphicValue:question;" json:"property,omitempty"`
	Entries      []Entry       `gorm:"polymorphic:Owner;polymorphicValue:question;" json:"entries,omitempty"`
}
