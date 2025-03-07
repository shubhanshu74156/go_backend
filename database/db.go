// database/connect.go

package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shubhanshu74156/go-rest-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the database connection and runs migrations.
func Connect() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Formulate DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open connection to database
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Run migrations
	err = DB.AutoMigrate(&models.Organization{}, &models.User{}, &models.Artist{}, &models.Album{}, &models.Track{}, &models.Favorite{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database connected and migrated successfully!")
}
