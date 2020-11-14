package model

type Answer struct {
	Model
	UserID   uint `gorm:"not null;" json:"-"`
	ChoiceID uint `gorm:"not null;" json:"-"`
}
