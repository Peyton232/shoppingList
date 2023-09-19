// router.go

package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router and sets up the routes for the shopping list API.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Use the middleware for all routes
	r.Use(CheckUserToken())

	// Group routes under /api
	api := r.Group("/api")

	// Routes for shopping list
	api.POST("/add-item", AddItemHandler)
	api.DELETE("/remove-item/:itemID", RemoveItemHandler)
	api.PUT("/update-item/:itemID", UpdateItemHandler)
	api.GET("/get-shopping-list", GetShoppingListHandler)

	return r
}
