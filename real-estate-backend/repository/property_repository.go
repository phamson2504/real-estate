package repository

import "real-estate-backend/model"

type PropertyRepository interface {
	Save(property *model.Property)
	Update(property model.Property)
	Delete(propertyId int)
	FindAll() []model.Property
	FindByOffset(offset int, limit int) []model.Property
	TotalProperties() int64
	FindById(propertyId int) model.Property
	FindByAgentId(agentId int) []model.Property
	FindByBought(userId int) []model.Property
	FindBySold(userId int) []model.Property
}
