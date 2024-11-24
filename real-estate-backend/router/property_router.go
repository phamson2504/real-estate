package router

import (
	"real-estate-backend/controller"
	"real-estate-backend/middleware"
	"real-estate-backend/repository"

	"github.com/gin-gonic/gin"
)

func PropertyRouter(router *gin.RouterGroup, userRepository repository.UserRepository, PropertyController *controller.PropertyController) {
	propertyRouter := router.Group("/properties")
	propertyRouter.GET("", PropertyController.FindByPages)
	propertyRouter.GET("/all-properties", PropertyController.GetAllProperties)
	propertyRouter.GET("/properties-by-agent", PropertyController.GetPropertyByAgentId)
	propertyRouter.GET("/property-details", PropertyController.GetPropertyById)
	propertyRouter.POST("/create", middleware.DeserializeUser(userRepository), PropertyController.Create)
	propertyRouter.GET("/GetPropertyFavoreat", middleware.DeserializeUser(userRepository), PropertyController.GetPropertyFavoreat)
	propertyRouter.GET("/checkPropertyFavorite", middleware.DeserializeUser(userRepository), PropertyController.CheckPropertyFavoreat)
	propertyRouter.GET("/propertyFavoriteByUserId", middleware.DeserializeUser(userRepository), PropertyController.PropertyFavoriteByUserId)
}
