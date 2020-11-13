package model

import (
	"gorm.io/gorm"
)

type Choice struct {
	gorm.Model
	VerifiedAt NullTime
	Title      string   `gorm:"size:255;not null;"`
	QuestionID uint     `gorm:"not null;" json:"-"`
	Users      []User   `gorm:"many2many:answers;" json:",omitempty"`
	Answers    []Answer `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"` // constraint
}
