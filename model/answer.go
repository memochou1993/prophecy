package model

import (
	"gorm.io/gorm"
	"time"
)

type Answer struct {
	gorm.Model
	VerifiedAt time.Time
	Title      string `gorm:"size:255;not null;"`
	QuestionID uint   `gorm:"not null;"`
}
