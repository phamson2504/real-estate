package main

import (
	"log"
	"net/http"
	"os"
	"real-estate-backend/config"
	"real-estate-backend/controller"
	"real-estate-backend/helper"
	"real-estate-backend/repository"
	"real-estate-backend/router"
	"real-estate-backend/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	gin.SetMode(gin.ReleaseMode)

	// Create a folder to save the file if it does not exist
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		os.Mkdir("./uploads", 0755)
	}

	// database
	db := config.DatabaseConnection()
	validate := validator.New()

	// repository
	userRepository := repository.NewUserRepository(db)
	propertyRepository := repository.NewPropertyRepository(db)
	agentRepository := repository.NewAgentReposiotryImpl(db)
	imageRepository := repository.NewImageRepositoryImpl(db)
	transactionRepository := repository.NewTransactionRepositoryImpl(db)
	favorateRepsitory := repository.NewFavorateRepsitoryImpl(db)

	// service
	userService := service.NewUserServiceImpl(userRepository, validate)
	authenticationService := service.NewAuthenticationServiceImpl(userRepository, agentRepository, validate)
	propertyService := service.NewPropertyServiceImpl(propertyRepository, agentRepository, imageRepository, userRepository, validate)
	transactionServie := service.NewTransactionServiceImpl(transactionRepository, agentRepository, propertyRepository, imageRepository, userRepository)
	favorateService := service.NewFavorateServiceImpl(favorateRepsitory, agentRepository, propertyRepository, imageRepository)

	//Init controller
	authenticationController := controller.NewAuthenticationController(authenticationService, agentRepository)
	userController := controller.NewUserController(userService)
	propertyController := controller.NewPropertyController(propertyService, favorateService)
	transactionController := controller.NewTransactionController(transactionServie, favorateService)
	// router
	routes := router.NewRouter(userRepository, authenticationController, userController, propertyController, transactionController)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.PanicIfError(server_err)
}
