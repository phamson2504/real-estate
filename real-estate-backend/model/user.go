package model

type User struct {
	Id           int    `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"type:varchar(100);not null"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
	Email        string `gorm:"type:varchar(100);unique"`
	Role         string `gorm:"type:varchar(50);not null"`

	Agent     Agent      `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	Favorites []Favorite `gorm:"foreignKey:UserId"`
	Reviews   []Review   `gorm:"foreignKey:UserId"`
}
