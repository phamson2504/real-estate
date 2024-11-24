package controller

import (
	"net/http"
	"real-estate-backend/data/request"
	"real-estate-backend/data/response"
	"real-estate-backend/helper"
	"real-estate-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TransactionController struct {
	TransactionService service.TransactionService
	FavorateService    service.FavorateService
}

func NewTransactionController(
	transactionService service.TransactionService,
	favorateService service.FavorateService,
) *TransactionController {
	return &TransactionController{
		TransactionService: transactionService,
		FavorateService:    favorateService,
	}
}

func (controller *TransactionController) Create(ctx *gin.Context) {
	log.Info().Msg("create transaction")

	TransactionCreateRequest := request.TransactionCreateRequest{}
	err := ctx.ShouldBindJSON(&TransactionCreateRequest)
	helper.PanicIfError(err)

	controller.TransactionService.Create(TransactionCreateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TransactionController) GetTransactionByUser(ctx *gin.Context) {
	log.Info().Msg("Get Property By User Id")
	userId := ctx.DefaultQuery("id", "0")
	id, err := strconv.Atoi(userId)
	if err != nil || id <= 0 {
		helper.PanicIfError(err)
	}

	transactions := controller.TransactionService.GetOfferdByUser(id)

	ctx.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}

func (controller *TransactionController) GetTransactionForSeller(ctx *gin.Context) {
	log.Info().Msg("Get Transaction For Seller")
	userId := ctx.DefaultQuery("id", "0")
	id, err := strconv.Atoi(userId)
	if err != nil || id <= 0 {
		helper.PanicIfError(err)
	}

	transactions := controller.TransactionService.GetOfferdForSeller(id)

	ctx.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}

func (controller *TransactionController) RequestTransactionForSeller(ctx *gin.Context) {
	log.Info().Msg("Request Transaction For Seller")

	var data map[string]interface{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	transactionId := int(data["id"].(float64))
	status := int(data["status"].(float64))

	err := controller.TransactionService.UpdateRequest(transactionId, status)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Request processed successfully",
	})
}

func (controller *TransactionController) GettransactionSold(ctx *gin.Context) {
	log.Info().Msg("Get Properties Sold")
	userId := ctx.DefaultQuery("id", "0")
	id, err := strconv.Atoi(userId)
	if err != nil || id <= 0 {
		helper.PanicIfError(err)
	}

	transactions := controller.TransactionService.GetTransactionSold(id)

	ctx.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}

func (controller *TransactionController) GetTransactionBougth(ctx *gin.Context) {
	log.Info().Msg("Get Transaction Bougth")
	userId := ctx.DefaultQuery("id", "0")
	id, err := strconv.Atoi(userId)
	if err != nil || id <= 0 {
		helper.PanicIfError(err)
	}

	transactions := controller.TransactionService.GetTransactionBought(id)

	ctx.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}
