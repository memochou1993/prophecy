package model

type House struct {
	Model
	Name      string     `gorm:"size:255;not null;uniqueIndex;" json:"name" validate:"required"`
	Settings  string     `gorm:"type:json;default:null" json:"settings"`
	OwnerID   uint       `gorm:"foreignKey:OwnerID;not null;" json:"ownerID"`
	Owner     *User      `gorm:"constraint:OnDelete:CASCADE;" json:"owner,omitempty"`
	Users     []User     `gorm:"many2many:members;" json:"users,omitempty"`
	Members   []Member   `gorm:"constraint:OnDelete:CASCADE;" json:"-"` // constraint
	Questions []Question `gorm:"constraint:OnDelete:CASCADE;" json:"questions,omitempty"`
	Point     *Point     `gorm:"constraint:OnDelete:CASCADE;" json:"point,omitempty"`
	Property  *Property  `gorm:"polymorphic:Owner;polymorphicValue:house;" json:"property,omitempty"`
	Entries   []Entry    `gorm:"polymorphic:Owner;polymorphicValue:house;" json:"entries,omitempty"`
}
