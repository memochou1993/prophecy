package model

type Property struct {
	Model
	Amount    int64  `gorm:"not null;" json:"amount"`
	PointID   uint   `gorm:"not null;" json:"pointID"`
	OwnerID   uint   `json:"-"`
	OwnerType string `gorm:"size:20;not null;" json:"-"`
	Point     Point  `json:"point,omitempty"`
}
