package router

import (
	"real-estate-backend/controller"
	"real-estate-backend/repository"

	"github.com/gin-gonic/gin"
)

func TransactionRouter(router *gin.RouterGroup, userRepository repository.UserRepository, transactionController *controller.TransactionController) {
	transactionRouter := router.Group("/transaction")
	transactionRouter.POST("", transactionController.Create)
	transactionRouter.GET("/get-properties-offered", transactionController.GetTransactionByUser)
	transactionRouter.GET("/get-transaction-offered-seller", transactionController.GetTransactionForSeller)
	transactionRouter.POST("/request-transaction-for-seller", transactionController.RequestTransactionForSeller)
	transactionRouter.GET("/getTransactionSold", transactionController.GettransactionSold)
	transactionRouter.GET("/getTransactionBougth", transactionController.GetTransactionBougth)
}
