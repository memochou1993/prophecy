package model

import "gorm.io/gorm"

type Participant struct {
	gorm.Model
	UserID     uint `gorm:"not null;" json:"-"`
	QuestionID uint `gorm:"not null;" json:"-"`
	IsOwner    bool `gorm:"not null;"`
}
