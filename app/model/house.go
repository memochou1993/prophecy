package model

import "gorm.io/gorm"

type House struct {
	gorm.Model
	Name      string     `gorm:"size:255;not null;uniqueIndex;"`
	Settings  string     `gorm:"type:json;"`
	UserID    uint       `gorm:"not null;" json:"-"`
	Point     *Point     `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"`
	Questions []Question `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"`
	Users     []User     `gorm:"many2many:house;" json:",omitempty"`
	Members   []Member   `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"` // constraint
	Property  *Property  `gorm:"polymorphic:Owner;polymorphicValue:house;" json:",omitempty"`
	Entries   []Entry    `gorm:"polymorphic:Owner;polymorphicValue:house;" json:",omitempty"`
}
