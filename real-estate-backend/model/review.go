package model

type Review struct {
	Id         int `gorm:"primaryKey;autoIncrement"`
	UserId     int
	PropertyId int
	Rating     int
	ReviewText string
	Status     int

	User     User     `gorm:"foreignKey:UserId"`
	Property Property `gorm:"foreignKey:PropertyId"`
}
