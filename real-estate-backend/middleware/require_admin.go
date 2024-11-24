package middleware

import (
	"net/http"
	"real-estate-backend/model"

	"github.com/gin-gonic/gin"
)

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("currentUser")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if usr, ok := user.(*model.User); ok {
			if usr.Role != "admin" {
				c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admins only"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
			c.Abort()
			return
		}

		c.Next() // Proceed if the user is an admin
	}
}
