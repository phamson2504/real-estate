package repository

import (
	"errors"
	"real-estate-backend/helper"
	"real-estate-backend/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// FindByEmail implements UserRepository.
func (u *UserRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// FindByNameOrEmail implements UserRepository.
func (u *UserRepositoryImpl) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := u.Db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var user model.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.PanicIfError(result.Error)
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() []model.User {
	var users []model.User
	result := u.Db.Find(&users)
	helper.PanicIfError(result.Error)

	return users
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(userId int) (*model.User, error) {
	var user model.User
	result := u.Db.Find(&user, userId)
	if result != nil {
		return &user, nil
	} else {
		return &user, errors.New("user is not found")
	}
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(user model.User) {
	result := u.Db.Create(&user)
	helper.PanicIfError(result.Error)
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user model.User) {
	result := u.Db.Model(&user).Updates(user)
	helper.PanicIfError(result.Error)
}
