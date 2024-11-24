package model

import "time"

type Property struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	Title       string
	Description string
	MinPrice    float64 `gorm:"not null"`
	MaxPrice    float64 `gorm:"not null"`
	Location    string
	Bedrooms    int
	Bathrooms   int
	SquareFeet  int
	AgentId     int // Foreign key to Agent
	Status      int
	CreatedAt   time.Time
	CityId      int
	DistrictId  int

	Agent        Agent         `gorm:"foreignKey:AgentId"`
	Images       []Image       `gorm:"foreignKey:PropertyId;constraint:OnDelete:CASCADE;"`
	Transactions []Transaction `gorm:"foreignKey:PropertyId"`
	Favorites    []Favorite    `gorm:"foreignKey:PropertyId;constraint:OnDelete:CASCADE;"`
	Reviews      []Review      `gorm:"foreignKey:PropertyId;constraint:OnDelete:CASCADE;"`
}
