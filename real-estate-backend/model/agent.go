package model

type Agent struct {
	Id            int `gorm:"primaryKey;autoIncrement"`
	UserId        int `gorm:"uniqueIndex"`
	AgencyName    string
	ContactNumber string
	AvatarAgent   string

	User       *User      `gorm:"constraint:OnDelete:CASCADE;"`
	Properties []Property `gorm:"foreignKey:AgentId"`
}
