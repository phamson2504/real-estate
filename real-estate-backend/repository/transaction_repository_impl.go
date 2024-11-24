package repository

import (
	"errors"
	"fmt"
	"real-estate-backend/helper"
	"real-estate-backend/model"

	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	Db *gorm.DB
}

func NewTransactionRepositoryImpl(Db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{Db: Db}
}

// save implements TransactionRepository.
func (t *TransactionRepositoryImpl) Save(transaction *model.Transaction) {
	result := t.Db.Create(transaction)
	helper.PanicIfError(result.Error)
}

// GetBySellerId implements TransactionRepository.
func (t *TransactionRepositoryImpl) GetBySellerId(sellerId int) ([]model.Transaction, error) {
	transaction := []model.Transaction{}
	err := t.Db.Model(&model.Transaction{}).Where("seller_id =?", sellerId).Order("status ASC").Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// GetByUserId implements TransactionRepository.
func (t *TransactionRepositoryImpl) GetByUserId(userId int) ([]model.Transaction, error) {
	transaction := []model.Transaction{}
	err := t.Db.Model(&model.Transaction{}).Where("buyer_id =?", userId).Order("status ASC").Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// UpdateRequest implements TransactionRepository.
func (t *TransactionRepositoryImpl) UpdateRequest(transactionId int, status int) error {
	result := t.Db.Model(&model.Transaction{}).
		Where("id = ?", transactionId).
		Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no transaction found with ID %d", transactionId)
	}
	return nil
}

// CheckPropertySold implements TransactionRepository.
func (t *TransactionRepositoryImpl) CheckPropertySold(propertyId int) bool {
	var transaction model.Transaction
	result := t.Db.Where("property_id = ? AND status != 2", propertyId).First(&transaction)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}

// GetTransactionSold implements TransactionRepository.
func (t *TransactionRepositoryImpl) GetTransactionSold(userId int) ([]model.Transaction, error) {
	var transaction []model.Transaction
	err := t.Db.Model(&model.Transaction{}).
		Where("status = 2 AND seller_id = ?", userId).
		Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// GetTransactionBought implements TransactionRepository.
func (t *TransactionRepositoryImpl) GetTransactionBought(userId int) ([]model.Transaction, error) {
	var transaction []model.Transaction
	err := t.Db.Model(&model.Transaction{}).
		Where("status = 2 AND buyer_id = ?", userId).
		Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
