package router

import (
	"real-estate-backend/controller"
	"real-estate-backend/middleware"
	"real-estate-backend/repository"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup, userRepository repository.UserRepository, userController *controller.UserController) {
	userRouter := router.Group("/users")
	userRouter.POST("", userController.Create)
	userRouter.DELETE("/:userId", userController.Delete)
	userRouter.PATCH("/:userId", userController.Update)
	userRouter.GET("/:userId", userController.FindById)
	userRouter.GET("", middleware.DeserializeUser(userRepository), userController.FindAll)
}
