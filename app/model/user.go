package model

import (
	"encoding/json"
)

type User struct {
	Model
	Name            string        `gorm:"size:255;not null;" json:"name"`
	Email           string        `gorm:"size:255;not null;uniqueIndex;" json:"email"`
	Password        string        `gorm:"size:255;not null;" json:"password,omitempty"`
	OwnedHouses     []House       `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE;" json:"ownedHouses,omitempty"`
	JoinedHouses    []House       `gorm:"many2many:members;" json:"joinedHouses,omitempty"`
	Members         []Member      `gorm:"constraint:OnDelete:CASCADE;" json:"-"` // constraint
	OwnedQuestions  []Question    `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE;" json:"ownedQuestions,omitempty"`
	JoinedQuestions []Question    `gorm:"many2many:participants;" json:"joinedQuestions,omitempty"`
	Participants    []Participant `gorm:"constraint:OnDelete:CASCADE;" json:"-"` // constraint
	Choices         []Choice      `gorm:"many2many:answers;" json:"choices,omitempty"`
	Answers         []Answer      `gorm:"constraint:OnDelete:CASCADE;" json:"-"` // constraint
	Properties      []Property    `gorm:"polymorphic:Owner;polymorphicValue:user;" json:"properties,omitempty"`
	Entries         []Entry       `gorm:"polymorphic:Owner;polymorphicValue:user;" json:"entries,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type user User

	item := user(u)
	item.Password = ""

	return json.Marshal(item)
}
