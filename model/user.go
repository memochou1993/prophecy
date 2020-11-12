package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string        `gorm:"size:255;not null;"`
	Email        string        `gorm:"size:255;not null;uniqueIndex;"`
	Password     string        `gorm:"size:255;not null;"`
	House        House         `gorm:"constraint:OnDelete:CASCADE;"`
	Answers      []Answer      `gorm:"constraint:OnDelete:CASCADE;"`
	Participants []Participant `gorm:"constraint:OnDelete:CASCADE;"`
	Property     Property      `gorm:"polymorphic:Owner;polymorphicValue:user;"`
	Entries      []Entry       `gorm:"polymorphic:Owner;polymorphicValue:user;"`
}
