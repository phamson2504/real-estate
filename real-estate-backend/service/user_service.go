package service

import (
	"real-estate-backend/data/request"
	"real-estate-backend/data/response"
)

type UserService interface {
	Create(userReq request.UserCreateRequest)
	Update(userReq request.UserUpdateRequest)
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindAll() []response.UserResponse
}
