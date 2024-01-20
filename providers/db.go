package providers

import (
	"fmt"
	"os"

	"github.com/aaqyaar/telemedicine/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

/**
 * @title       : db.go
 * @description : database connection
 * @package     : providers
 */
func Connect() {

	dsn := os.Getenv("DATABASE_URL")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database Connected")

	// Migrate the schema
	AutoMigrate(DB)
}

func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return DB
}
