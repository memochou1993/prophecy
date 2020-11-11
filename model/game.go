package model

import (
	"gorm.io/gorm"
	"time"
)

type Game struct {
	gorm.Model
	ReviewedAt  time.Time
	PublishedAt time.Time
	OpenedAt    time.Time
	ClosedAt    time.Time
	Name        string `gorm:"size:255;not null"`
	UserID      uint
}
