package main

import (
	"log"
	"shoppingCartAPI/intenal/api"
)

func main() {
	router := api.SetupRouter()

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
