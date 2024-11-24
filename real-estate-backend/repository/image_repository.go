package repository

import "real-estate-backend/model"

type ImageRepository interface {
	Save(image model.Image)
	Update(image model.Image)
	Delete(imageId int)
	FindByPropertyID(protertyId int) []model.Image
}
