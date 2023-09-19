package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the machine token from the request header
		userToken := c.GetHeader("X-User-Token")

		// Check if the user token is empty or not provided
		if userToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User token not provided"})
			c.Abort() // Abort the request, no further processing
			return
		}

		// Store the user token in the context for later use
		c.Set("UserToken", userToken)

		// Continue to the next handler
		c.Next()
	}
}
