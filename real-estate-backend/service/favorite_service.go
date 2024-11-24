package service

import (
	"real-estate-backend/data/request"
	"real-estate-backend/data/response"
)

type FavorateService interface {
	Create(favorateReq request.FavorateCreateRequest)
	CheckPropertyFavorite(favorateReq request.FavorateCreateRequest) bool
	PropertyFavoriteByUserId(userId int) []response.PropertyResponse
}
