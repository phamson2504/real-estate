package service

import (
	"fmt"
	"log"
	"real-estate-backend/data/request"
	"real-estate-backend/data/response"
	"real-estate-backend/helper"
	"real-estate-backend/model"
	"real-estate-backend/repository"
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	AgentRepository       repository.AgentRepository
	PropertyRepository    repository.PropertyRepository
	ImageRepository       repository.ImageRepository
	UserRepository        repository.UserRepository
}

func NewTransactionServiceImpl(
	transactionRepository repository.TransactionRepository,
	agentRepository repository.AgentRepository,
	propertyRepository repository.PropertyRepository,
	imageRepository repository.ImageRepository,
	userRepository repository.UserRepository,
) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		AgentRepository:       agentRepository,
		PropertyRepository:    propertyRepository,
		ImageRepository:       imageRepository,
		UserRepository:        userRepository,
	}

}

// Create implements TransactionService.
func (t *TransactionServiceImpl) Create(tranReq request.TransactionCreateRequest) {
	transaction := model.Transaction{
		PropertyId: tranReq.PropertyId,
		BuyerId:    tranReq.BuyerId,
		SellerId:   tranReq.SellerId,
		DateOffer:  tranReq.DateOffer,
		Amount:     tranReq.Amount,
		Status:     0,
	}
	t.TransactionRepository.Save(&transaction)
}

// GetOfferdByUser implements TransactionService.
func (t *TransactionServiceImpl) GetOfferdByUser(buyerId int) []response.TransactionResponse {
	transactions, err := t.TransactionRepository.GetByUserId(buyerId)
	helper.PanicIfError(err)

	baseURL := "http://localhost:8080"
	transactionsResp := []response.TransactionResponse{}

	for _, transaction := range transactions {
		property := t.PropertyRepository.FindById(transaction.PropertyId)
		agent := t.AgentRepository.FindByAgentId(property.AgentId)
		var images = t.ImageRepository.FindByPropertyID(property.Id)

		var avatarUrl = ""
		if agent.AvatarAgent != "" {
			avatarUrl = fmt.Sprintf("%s/%s", baseURL, agent.AvatarAgent)
		}

		agentResp := response.AgentResponse{
			Id:          agent.Id,
			Name:        agent.AgencyName,
			Contact:     agent.ContactNumber,
			AvatarAgent: avatarUrl,
		}

		propertyResp := response.PropertyResponse{
			Id:          property.Id,
			Title:       property.Title,
			Description: property.Description,
			MaxPrice:    property.MaxPrice,
			MinPrice:    property.MaxPrice,
			Location:    property.Location,
			Bedrooms:    property.Bedrooms,
			Bathrooms:   property.Bathrooms,
			SquareFeet:  property.SquareFeet,
			Agent:       agentResp,
		}

		if len(images) > 0 {
			propertyResp.Images = make([]response.ImageResponse, 1)
			var imageUrl = ""
			if images[0].URL != "" {
				imageUrl = fmt.Sprintf("%s/%s", baseURL, images[0].URL)
			}

			log.Print(propertyResp.Images[0].Url)
			propertyResp.Images[0] = response.ImageResponse{
				Id:  images[0].Id,
				Url: imageUrl,
			}
		}
		var statusTran string
		if transaction.Status == 0 {
			statusTran = "spending"
		} else if transaction.Status == 1 {
			statusTran = "accepted"
		} else {
			statusTran = "declined"
		}

		transactionResp := response.TransactionResponse{
			Id:         transaction.Id,
			Properties: propertyResp,
			Amount:     transaction.Amount,
			DateOffer:  transaction.DateOffer,
			Status:     statusTran,
		}

		transactionsResp = append(transactionsResp, transactionResp)
	}

	return transactionsResp
}

// GetOfferdForSeller implements TransactionService.
func (t *TransactionServiceImpl) GetOfferdForSeller(sellerId int) []response.TransactionResponse {
	transactions, err := t.TransactionRepository.GetBySellerId(sellerId)
	helper.PanicIfError(err)

	transactionsResp := []response.TransactionResponse{}

	for _, transaction := range transactions {
		property := t.PropertyRepository.FindById(transaction.PropertyId)
		agent := t.AgentRepository.FindByAgentId(property.AgentId)
		user, _ := t.UserRepository.FindById(agent.UserId)

		agentResp := response.AgentResponse{
			Id:      agent.Id,
			Email:   user.Email,
			Name:    agent.AgencyName,
			Contact: agent.ContactNumber,
		}

		propertyResp := response.PropertyResponse{
			Id:          property.Id,
			Title:       property.Title,
			Description: property.Description,
			MaxPrice:    property.MaxPrice,
			MinPrice:    property.MaxPrice,
			Location:    property.Location,
			Bedrooms:    property.Bedrooms,
			Bathrooms:   property.Bathrooms,
			SquareFeet:  property.SquareFeet,
			Agent:       agentResp,
		}

		var statusTran string
		if transaction.Status == 0 {
			statusTran = "spending"
		} else if transaction.Status == 1 {
			statusTran = "accepted"
		} else {
			statusTran = "declined"
		}

		transactionResp := response.TransactionResponse{
			Id:         transaction.Id,
			Properties: propertyResp,
			Amount:     transaction.Amount,
			DateOffer:  transaction.DateOffer,
			Status:     statusTran,
		}

		transactionsResp = append(transactionsResp, transactionResp)
	}

	return transactionsResp
}

// UpdateRequest implements TransactionService.
func (t *TransactionServiceImpl) UpdateRequest(transactionId int, status int) error {
	return t.TransactionRepository.UpdateRequest(transactionId, status)
}

// GetTransactionSold implements TransactionService.
func (t *TransactionServiceImpl) GetTransactionSold(userId int) []response.TransactionResponse {
	transactions, err := t.TransactionRepository.GetTransactionSold(userId)
	helper.PanicIfError(err)

	transactionsResp := []response.TransactionResponse{}
	baseURL := "http://localhost:8080"
	for _, transaction := range transactions {
		property := t.PropertyRepository.FindById(transaction.PropertyId)
		agent := t.AgentRepository.FindByAgentId(property.AgentId)
		user, _ := t.UserRepository.FindById(agent.UserId)

		var avatarUrl = ""
		if agent.AvatarAgent != "" {
			avatarUrl = fmt.Sprintf("%s/%s", baseURL, agent.AvatarAgent)
		}

		agentResp := response.AgentResponse{
			Id:          agent.Id,
			Email:       user.Email,
			Name:        agent.AgencyName,
			Contact:     agent.ContactNumber,
			AvatarAgent: avatarUrl,
		}

		propertyResp := response.PropertyResponse{
			Id:          property.Id,
			Title:       property.Title,
			Description: property.Description,
			MaxPrice:    property.MaxPrice,
			MinPrice:    property.MaxPrice,
			Location:    property.Location,
			Bedrooms:    property.Bedrooms,
			Bathrooms:   property.Bathrooms,
			SquareFeet:  property.SquareFeet,
			Agent:       agentResp,
		}

		var statusTran string
		if transaction.Status == 0 {
			statusTran = "spending"
		} else if transaction.Status == 1 {
			statusTran = "accepted"
		} else {
			statusTran = "declined"
		}

		transactionResp := response.TransactionResponse{
			Id:         transaction.Id,
			Properties: propertyResp,
			Amount:     transaction.Amount,
			DateOffer:  transaction.DateOffer,
			Status:     statusTran,
		}

		transactionsResp = append(transactionsResp, transactionResp)
	}
	return transactionsResp
}

// GetTransactionBought implements TransactionService.
func (t *TransactionServiceImpl) GetTransactionBought(userId int) []response.TransactionResponse {
	transactions, err := t.TransactionRepository.GetTransactionBought(userId)
	helper.PanicIfError(err)

	transactionsResp := []response.TransactionResponse{}
	baseURL := "http://localhost:8080"
	for _, transaction := range transactions {
		property := t.PropertyRepository.FindById(transaction.PropertyId)
		agent := t.AgentRepository.FindByAgentId(property.AgentId)
		user, _ := t.UserRepository.FindById(agent.UserId)

		var avatarUrl = ""
		if agent.AvatarAgent != "" {
			avatarUrl = fmt.Sprintf("%s/%s", baseURL, agent.AvatarAgent)
		}

		agentResp := response.AgentResponse{
			Id:          agent.Id,
			Email:       user.Email,
			Name:        agent.AgencyName,
			Contact:     agent.ContactNumber,
			AvatarAgent: avatarUrl,
		}

		propertyResp := response.PropertyResponse{
			Id:          property.Id,
			Title:       property.Title,
			Description: property.Description,
			MaxPrice:    property.MaxPrice,
			MinPrice:    property.MaxPrice,
			Location:    property.Location,
			Bedrooms:    property.Bedrooms,
			Bathrooms:   property.Bathrooms,
			SquareFeet:  property.SquareFeet,
			Agent:       agentResp,
		}

		var statusTran string
		if transaction.Status == 0 {
			statusTran = "spending"
		} else if transaction.Status == 1 {
			statusTran = "accepted"
		} else {
			statusTran = "declined"
		}

		transactionResp := response.TransactionResponse{
			Id:         transaction.Id,
			Properties: propertyResp,
			Amount:     transaction.Amount,
			DateOffer:  transaction.DateOffer,
			Status:     statusTran,
		}

		transactionsResp = append(transactionsResp, transactionResp)
	}
	return transactionsResp
}
