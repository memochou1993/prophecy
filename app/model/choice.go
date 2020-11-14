package model

type Choice struct {
	Model
	VerifiedAt NullTime `json:"verifiedAt"`
	Title      string   `gorm:"size:255;not null;" json:"title"`
	QuestionID uint     `gorm:"not null;" json:"-" json:"questionID"`
	Users      []User   `gorm:"many2many:answers;" json:"users,omitempty"`
	Answers    []Answer `gorm:"constraint:OnDelete:CASCADE;" json:"-"` // constraint
}
