// main.go

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shubhanshu74156/go-rest-api/database"
	"github.com/shubhanshu74156/go-rest-api/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database and run migrations
	database.Connect()

	// Register routes
	router := routes.SetupRouter()

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not set
	}

	log.Printf("Server is running on port %s", port)
	router.Run(":" + port) // Start the server
}
