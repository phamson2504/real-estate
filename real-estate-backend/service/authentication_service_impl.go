package service

import (
	"errors"
	"real-estate-backend/config"
	"real-estate-backend/data/request"
	"real-estate-backend/helper"
	"real-estate-backend/model"
	"real-estate-backend/repository"
	"real-estate-backend/utils"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UserRepository  repository.UserRepository
	AgentRepository repository.AgentRepository
	Validate        *validator.Validate
}

func NewAuthenticationServiceImpl(userRepository repository.UserRepository, agentRepository repository.AgentRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UserRepository:  userRepository,
		AgentRepository: agentRepository,
		Validate:        validate,
	}
}

// Login implements AuthenticationService.
func (a *AuthenticationServiceImpl) Login(userReq request.LoginRequest) (string, error) {
	user, err := a.UserRepository.FindByUsername(userReq.Username)
	if err != nil {
		return "", errors.New("user not found")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(user.PasswordHash, userReq.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, user.Id, config.TokenSecret)
	helper.PanicIfError(err_token)
	return token, nil
}

// Register implements AuthenticationService.
func (a *AuthenticationServiceImpl) Register(userReq request.UserCreateRequest) error {
	err := a.Validate.Struct(userReq)
	helper.PanicIfError(err)

	hashedPassword, errPass := utils.HashPassword(userReq.Password)
	helper.PanicIfError(errPass)

	userByEmail, _ := a.UserRepository.FindByEmail(userReq.Email)
	if userByEmail != nil {
		return errors.New("email already exists")
	}

	userByUsername, _ := a.UserRepository.FindByUsername(userReq.Username)
	if userByUsername != nil {
		return errors.New("username already exists")
	}

	user := model.User{
		Username:     userReq.Username,
		PasswordHash: hashedPassword,
		Email:        userReq.Email,
		Role:         userReq.Role,
	}

	a.UserRepository.Save(user)
	return nil
}

// UpdateProfile implements AuthenticationService.
func (a *AuthenticationServiceImpl) UpdateProfile(agent model.Agent, email string, currentUser model.User) {
	if currentUser.Email != email {
		a.UserRepository.Update(currentUser)
	}
	if agent.Id == 0 {
		a.AgentRepository.Save(agent)
	} else {
		a.AgentRepository.Update(agent)
	}
}
