package repository

import (
	"real-estate-backend/helper"
	"real-estate-backend/model"

	"gorm.io/gorm"
)

type PropertyRepositoryImpl struct {
	Db *gorm.DB
}

func NewPropertyRepository(Db *gorm.DB) PropertyRepository {
	return &PropertyRepositoryImpl{Db: Db}
}

// Delete implements PropertyRepository.
func (p *PropertyRepositoryImpl) Delete(propertyId int) {
	var property model.Property
	result := p.Db.Where("id = ?", propertyId).Delete(&property)
	helper.PanicIfError(result.Error)
}

// Save implements PropertyRepository.
func (p *PropertyRepositoryImpl) Save(property *model.Property) {
	result := p.Db.Create(property)
	helper.PanicIfError(result.Error)
}

// Update implements PropertyRepository.
func (p *PropertyRepositoryImpl) Update(property model.Property) {
	result := p.Db.Model(&property).Updates(property)
	helper.PanicIfError(result.Error)
}

// FindAll implements PropertyRepository.
func (p *PropertyRepositoryImpl) FindAll() []model.Property {
	var properties []model.Property
	result := p.Db.Order("created_at ASC").Find(&properties)
	helper.PanicIfError(result.Error)

	return properties
}

// FindByOffset implements PropertyRepository.
func (p *PropertyRepositoryImpl) FindByOffset(offset int, limit int) []model.Property {
	var properties []model.Property
	result := p.Db.Model(&model.Property{}).
		Joins("LEFT JOIN transactions ON properties.id = transactions.property_id").
		Where("(transactions.id IS NULL OR transactions.status != 2) AND properties.status = 1").
		Group("properties.id").
		Order("properties.created_at ASC").
		Offset(offset).
		Limit(limit).
		Find(&properties)
	helper.PanicIfError(result.Error)

	return properties
}

// TotalProperties implements PropertyRepository.
func (p *PropertyRepositoryImpl) TotalProperties() int64 {
	var totalItems int64
	result := p.Db.Model(&model.Property{}).Count(&totalItems)
	helper.PanicIfError(result.Error)

	return totalItems
}

// FindById implements PropertyRepository.
func (p *PropertyRepositoryImpl) FindById(propertyId int) model.Property {
	var property model.Property
	result := p.Db.Model(&model.Property{}).Where("id = ?", propertyId).First(&property)
	helper.PanicIfError(result.Error)
	return property
}

// FindByAgentId implements PropertyRepository.
func (p *PropertyRepositoryImpl) FindByAgentId(agentId int) []model.Property {
	var property []model.Property
	result := p.Db.Model(&model.Property{}).Where("agent_id = ?", agentId).Find(&property)
	helper.PanicIfError(result.Error)
	return property
}

// FindByBought implements PropertyRepository.
func (p *PropertyRepositoryImpl) FindByBought(userId int) []model.Property {
	var properties []model.Property
	result := p.Db.Model(&model.Property{}).
		Joins("INNER JOIN transactions ON properties.id = transactions.property_id").
		Where("transactions.status = 2 AND transactions.buyer_id = ?", userId).
		Find(&properties)

	helper.PanicIfError(result.Error)
	return properties
}

// FindBySold implements PropertyRepository.
func (p *PropertyRepositoryImpl) FindBySold(userId int) []model.Property {
	var properties []model.Property
	result := p.Db.Model(&model.Property{}).
		Joins("INNER JOIN transactions ON properties.id = transactions.property_id").
		Where("transactions.status = 2 AND transactions.seller_id = ?", userId).
		Find(&properties)

	helper.PanicIfError(result.Error)
	return properties
}
