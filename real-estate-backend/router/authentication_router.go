package router

import (
	"net/http"
	"real-estate-backend/controller"
	"real-estate-backend/middleware"
	"real-estate-backend/repository"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	userRepository repository.UserRepository,
	authenticationController *controller.AuthenticationController,
	userController *controller.UserController,
	propertyController *controller.PropertyController,
	transactionController *controller.TransactionController,
) *gin.Engine {
	service := gin.Default()
	service.Use(cors.Default())

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router := service.Group("/api")
	authenticationRouter := router.Group("/authentication")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)
	authenticationRouter.GET("/user", middleware.DeserializeUser(userRepository), authenticationController.CurrentUser)
	authenticationRouter.POST("/update-profile", middleware.DeserializeUser(userRepository), authenticationController.UpdateProfile)

	UserRouter(router, userRepository, userController)
	PropertyRouter(router, userRepository, propertyController)
	TransactionRouter(router, userRepository, transactionController)

	service.Static("/uploads", "./uploads")

	return service
}
