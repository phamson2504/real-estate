package model

type Image struct {
	Id         int    `gorm:"primaryKey;autoIncrement"`
	PropertyId int    `gorm:"not null"` // Foreign key to Property table
	URL        string `gorm:"type:text;not null"`

	Property Property `gorm:"foreignKey:PropertyId"`
}
