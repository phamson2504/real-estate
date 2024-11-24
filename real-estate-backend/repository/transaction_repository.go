package repository

import (
	"real-estate-backend/model"
)

type TransactionRepository interface {
	Save(transaction *model.Transaction)
	GetByUserId(userId int) ([]model.Transaction, error)
	GetBySellerId(sellerId int) ([]model.Transaction, error)
	UpdateRequest(transactionId int, status int) error
	CheckPropertySold(propertyId int) bool
	GetTransactionSold(userId int) ([]model.Transaction, error)
	GetTransactionBought(userId int) ([]model.Transaction, error)
}
