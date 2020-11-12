package model

import (
	"gorm.io/gorm"
	"time"
)

type Choice struct {
	gorm.Model
	VerifiedAt time.Time
	Title      string   `gorm:"size:255;not null;"`
	QuestionID uint     `gorm:"not null;"`
	Answers    []Answer `gorm:"constraint:OnDelete:CASCADE;"`
}
