package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string        `gorm:"size:255;not null;"`
	Email          string        `gorm:"size:255;not null;uniqueIndex;"`
	Password       string        `gorm:"size:255;not null;" json:",omitempty"`
	OwnedHouses    []House       `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE;" json:",omitempty"`
	Houses         []House       `gorm:"many2many:members;" json:",omitempty"`
	Members        []Member      `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"` // constraint
	OwnedQuestions []Question    `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE;" json:",omitempty"`
	Questions      []Question    `gorm:"many2many:participants;" json:",omitempty"`
	Participants   []Participant `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"` // constraint
	Choices        []Choice      `gorm:"many2many:answers;" json:",omitempty"`
	Answers        []Answer      `gorm:"constraint:OnDelete:CASCADE;" json:",omitempty"` // constraint
	Properties     []Property    `gorm:"polymorphic:Owner;polymorphicValue:user;" json:",omitempty"`
	Entries        []Entry       `gorm:"polymorphic:Owner;polymorphicValue:user;" json:",omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type user User

	x := user(u)
	x.Password = ""

	return json.Marshal(x)
}
