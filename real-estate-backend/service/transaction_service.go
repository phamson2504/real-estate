package service

import (
	"real-estate-backend/data/request"
	"real-estate-backend/data/response"
)

type TransactionService interface {
	Create(tranReq request.TransactionCreateRequest)
	GetOfferdByUser(buyerId int) []response.TransactionResponse
	GetOfferdForSeller(buyerId int) []response.TransactionResponse
	UpdateRequest(transactionId int, status int) error
	GetTransactionSold(userId int) []response.TransactionResponse
	GetTransactionBought(userId int) []response.TransactionResponse
}
