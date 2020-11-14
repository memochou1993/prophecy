package model

type Participant struct {
	Model
	UserID     uint `gorm:"not null;" json:"-"`
	QuestionID uint `gorm:"not null;" json:"-"`
}
