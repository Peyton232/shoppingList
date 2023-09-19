// handler.go

package api

import (
	"fmt"
	"net/http"

	"shoppingCartAPI/intenal/models"

	"github.com/gin-gonic/gin"
)

// AddItemHandler handles the request to add an item to the shopping list.
func AddItemHandler(c *gin.Context) {
	// test
	fmt.Println(c.GetString("UserToken"))

	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// assign item to grocery list

	c.JSON(http.StatusOK, newItem)
}

// RemoveItemHandler handles the request to remove an item from the shopping list.
func RemoveItemHandler(c *gin.Context) {
	itemID, err := models.ParseItemID(c.Param("itemID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(itemID)

	// remove item with ID itemID from shopping list

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// UpdateItemHandler handles the request to update an item in the shopping list.
func UpdateItemHandler(c *gin.Context) {
	itemID, err := models.ParseItemID(c.Param("itemID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(itemID)

	var updatedItem models.Item

	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find item in hsopping list and update

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// GetShoppingListHandler handles the request to retrieve the shopping list.
func GetShoppingListHandler(c *gin.Context) {
	items := []models.Item{
		{
			ID:          1,
			Name:        "Item 1",
			Description: "Description for Item 1",
			Amount:      "5.99",
			Completed:   false,
			CreatedAt:   "2023-09-18T10:00:00Z",
			UpdatedAt:   "2023-09-18T11:30:00Z",
		},
		{
			ID:          2,
			Name:        "Item 2",
			Description: "Description for Item 2",
			Amount:      "12.50",
			Completed:   true,
			CreatedAt:   "2023-09-19T14:15:00Z",
			UpdatedAt:   "2023-09-19T15:45:00Z",
		},
		// Add more items as needed
	}

	c.JSON(http.StatusOK, items)
}
