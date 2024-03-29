package providers

import (
	"fmt"

	"github.com/aaqyaar/telemedicine/models"
	"github.com/aaqyaar/telemedicine/utils"
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

	dsn := utils.GetEnv("DATABASE_URL")

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Database Connected")

	// Migrate the schema
	AutoMigrate(DB)
}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&models.User{},
	)
}
