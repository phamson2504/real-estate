package repository

import (
	"real-estate-backend/model"
)

type UserRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(userId int)
	FindById(userId int) (*model.User, error)
	FindAll() []model.User
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}
