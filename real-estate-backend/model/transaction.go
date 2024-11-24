package model

type Transaction struct {
	Id         int `gorm:"primaryKey;autoIncrement"`
	PropertyId int
	BuyerId    int
	SellerId   int
	Amount     float64
	Status     int
	DateOffer  string

	Property Property `gorm:"foreignKey:PropertyId"`
	Buyer    User     `gorm:"foreignKey:BuyerId"`
	Seller   User     `gorm:"foreignKey:SellerId"`
}
