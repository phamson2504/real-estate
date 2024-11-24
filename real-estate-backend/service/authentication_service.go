package service

import (
	"real-estate-backend/data/request"
	"real-estate-backend/model"
)

type AuthenticationService interface {
	Login(userReq request.LoginRequest) (string, error)
	Register(userReq request.UserCreateRequest) error
	UpdateProfile(agentReq model.Agent, email string, currentUser model.User)
}
