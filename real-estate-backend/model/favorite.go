package model

type Favorite struct {
	Id         int `gorm:"primaryKey;autoIncrement"`
	UserId     int
	PropertyId int
	User       User     `gorm:"foreignKey:UserId"`
	Property   Property `gorm:"foreignKey:PropertyId"`
}
