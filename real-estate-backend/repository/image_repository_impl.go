package repository

import (
	"real-estate-backend/helper"
	"real-estate-backend/model"

	"gorm.io/gorm"
)

type ImageRepositoryImpl struct {
	Db *gorm.DB
}

func NewImageRepositoryImpl(Db *gorm.DB) ImageRepository {
	return &ImageRepositoryImpl{Db: Db}
}

// Delete implements ImageRepository.
func (i *ImageRepositoryImpl) Delete(imageId int) {
	panic("unimplemented")
}

// Save implements ImageRepository.
func (i *ImageRepositoryImpl) Save(image model.Image) {
	result := i.Db.Create(&image)
	helper.PanicIfError(result.Error)
}

// Update implements ImageRepository.
func (i *ImageRepositoryImpl) Update(image model.Image) {
	panic("unimplemented")
}

// findByPropertyId implements ImageRepository.
func (i *ImageRepositoryImpl) FindByPropertyID(protertyId int) []model.Image {
	var images []model.Image
	result := i.Db.Where("property_id = ?", protertyId).Find(&images)
	helper.PanicIfError(result.Error)

	return images
}
