package model

import (
	"gorm.io/gorm"
	"time"
)

type Question struct {
	gorm.Model
	ReviewedAt  time.Time
	PublishedAt time.Time
	OpenedAt    time.Time
	ClosedAt    time.Time
	Title       string `gorm:"size:255;not null;"`
	Description string
	GameID      uint     `gorm:"not null;"`
	Choices     []Choice `gorm:"constraint:OnDelete:CASCADE;"`
}
