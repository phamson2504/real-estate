package model

type City struct {
	Id     int    `gorm:"primaryKey"`
	Name   string `gorm:"not null"`
	Status int

	Districts []District `gorm:"foreignKey:CityId;constraint:OnDelete:CASCADE;"`
}
