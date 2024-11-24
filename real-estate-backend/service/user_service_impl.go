package service

import (
	"real-estate-backend/data/request"
	"real-estate-backend/data/response"
	"real-estate-backend/helper"
	"real-estate-backend/model"
	"real-estate-backend/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, Validate: validate}
}

// Create implements UserService.
func (u *UserServiceImpl) Create(userReq request.UserCreateRequest) {
	err := u.Validate.Struct(userReq)
	helper.PanicIfError(err)

	hashedPassword, errPass := HashPassword(userReq.Password)
	helper.PanicIfError(errPass)

	user := model.User{
		Username:     userReq.Username,
		PasswordHash: hashedPassword,
		Email:        userReq.Email,
		Role:         userReq.Role,
	}
	u.UserRepository.Save(user)
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(userId int) {
	user, err := u.UserRepository.FindById(userId)
	helper.PanicIfError(err)
	u.UserRepository.Delete(user.Id)
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll() []response.UserResponse {
	users := u.UserRepository.FindAll()
	var userResp []response.UserResponse
	for _, value := range users {
		userResponse := response.UserResponse{
			Id:       value.Id,
			Username: value.Username,
			Email:    value.Email,
			Role:     value.Role,
		}
		userResp = append(userResp, userResponse)
	}
	return userResp
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById(userId int) response.UserResponse {
	user, err := u.UserRepository.FindById(userId)
	helper.PanicIfError(err)

	userResponse := response.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
	return response.UserResponse(userResponse)
}

// Update implements UserService.
func (u *UserServiceImpl) Update(userReq request.UserUpdateRequest) {
	var user, err = u.UserRepository.FindById(userReq.Id)
	helper.PanicIfError(err)

	user = &model.User{
		Id:           userReq.Id,
		Username:     userReq.Username,
		PasswordHash: userReq.Password,
		Email:        userReq.Email,
		Role:         userReq.Role,
	}
	u.UserRepository.Update(*user)
}

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
