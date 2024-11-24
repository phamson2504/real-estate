package service

import (
	"real-estate-backend/data/request"
	"real-estate-backend/data/response"
	"real-estate-backend/model"
)

type PropertyService interface {
	Create(user *model.User, property request.PropertyCreateRequest) error
	Update(user *model.User, property request.PropertyUpdateRequest) error
	FindByPages(page int, limit int) ([]response.PropertyResponse, int64, int)
	FindAll() []response.PropertyResponse
	FindById(propertyId int) response.PropertyResponse
	FindByAgentId(agentId int) []response.PropertyResponse
}
