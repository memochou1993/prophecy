package model

import "gorm.io/gorm"

type House struct {
	gorm.Model
	Name      string `gorm:"size:255;not null;uniqueIndex;"`
	Settings  string `gorm:"type:json;"`
	OwnerID   uint   `gorm:"foreignKey:OwnerID;not null;" json:"-"`
	Owner     User
	Users     []User     `gorm:"many2many:members;" json:",omitempty"`
	Members   []Member   `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"` // constraint
	Questions []Question `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"`
	Point     *Point     `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"`
	Property  *Property  `gorm:"polymorphic:Owner;polymorphicValue:house;" json:",omitempty"`
	Entries   []Entry    `gorm:"polymorphic:Owner;polymorphicValue:house;" json:",omitempty"`
}
