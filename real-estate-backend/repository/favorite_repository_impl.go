package repository

import (
	"errors"
	"real-estate-backend/helper"
	"real-estate-backend/model"

	"gorm.io/gorm"
)

type FavorateRepsitoryImpl struct {
	Db *gorm.DB
}

func NewFavorateRepsitoryImpl(Db *gorm.DB) FavorateRepsitory {
	return &FavorateRepsitoryImpl{Db: Db}
}

// Save implements FavorateRepsitory.
func (f *FavorateRepsitoryImpl) Save(favorate model.Favorite) {
	result := f.Db.Create(&favorate)
	helper.PanicIfError(result.Error)
}

// CheckPropertyFavorite implements FavorateRepsitory.
func (f *FavorateRepsitoryImpl) CheckPropertyFavorite(favorate model.Favorite) bool {
	var favorite model.Favorite
	result := f.Db.Where("user_id = ? AND property_id = ?", favorate.UserId, favorate.PropertyId).First(&favorite)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}

// GetPropertyFavoriteByUserId implements FavorateRepsitory.
func (f *FavorateRepsitoryImpl) GetPropertyFavoriteByUserId(userId int) []model.Favorite {
	var favorite []model.Favorite
	result := f.Db.Where("user_id = ?", userId).Find(&favorite)
	if result.Error != nil {
		return nil
	}
	return favorite
}
