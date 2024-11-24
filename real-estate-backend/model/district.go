package model

type District struct {
	Id     int    `gorm:"primaryKey"`
	Name   string `gorm:"not null"`
	CityId int    `gorm:"not null"`
	Status int

	City City `gorm:"foreignKey:CityId"`
}
