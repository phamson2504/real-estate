package repository

import "real-estate-backend/model"

type FavorateRepsitory interface {
	Save(favorate model.Favorite)
	CheckPropertyFavorite(favorate model.Favorite) bool
	GetPropertyFavoriteByUserId(userId int) []model.Favorite
}
