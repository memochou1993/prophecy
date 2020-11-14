package model

type Point struct {
	Model
	Name       string     `gorm:"size:255;not null;uniqueIndex;" json:"name"`
	HouseID    uint       `gorm:"not null;" json:"houseID"`
	Properties []Property `gorm:"constraint:OnDelete:CASCADE;" json:"properties,omitempty"`
	Entries    []Entry    `gorm:"constraint:OnDelete:CASCADE;" json:"entries,omitempty"`
}
