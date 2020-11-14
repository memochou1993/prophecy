package model

type Member struct {
	Model
	UserID  uint `gorm:"not null;" json:"-"`
	HouseID uint `gorm:"not null;" json:"-"`
}
