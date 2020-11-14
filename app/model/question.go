package model

type Question struct {
	Model
	ReviewedAt   NullTime      `json:"reviewedAt"`
	PublishedAt  NullTime      `json:"publishedAt"`
	OpenedAt     *Timestamp    `json:"openedAt" validate:"required"`
	ClosedAt     *Timestamp    `json:"closedAt" validate:"required"`
	Title        string        `gorm:"size:255;not null;" json:"title" validate:"required"`
	Description  string        `gorm:"size:255;not null;" json:"description"`
	HouseID      uint          `gorm:"not null;" json:"houseID"`
	OwnerID      uint          `gorm:"not null;" json:"ownerID"`
	Owner        *User         `gorm:"constraint:OnDelete:CASCADE;" json:"owner,omitempty"`
	Users        []User        `gorm:"many2many:participants;" json:"users,omitempty"`
	Participants []Participant `gorm:"constraint:OnDelete:CASCADE;" json:"-"` // constraint
	Choices      []Choice      `gorm:"constraint:OnDelete:CASCADE;" json:"choices,omitempty"`
	Property     *Property     `gorm:"polymorphic:Owner;polymorphicValue:question;" json:"property,omitempty"`
	Entries      []Entry       `gorm:"polymorphic:Owner;polymorphicValue:question;" json:"entries,omitempty"`
}
