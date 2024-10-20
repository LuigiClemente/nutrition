package main

import (
	"log"
	"nutrition/api"
	"nutrition/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// change to main
func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Set up the Gin router
	router := api.SetupRouter()

	go utils.ScheduleDeletion()
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)
	// Start the server
	router.Run(":1010")

}
