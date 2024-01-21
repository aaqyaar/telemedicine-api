package main

import (
	"github.com/aaqyaar/telemedicine/providers"
	"github.com/aaqyaar/telemedicine/routes/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

/**
 * @title       : main.go
 * @description : entry point of the application
 * @package     : main
 */
func main() {

	// load .env file
	godotenv.Load()

	// connect to database
	providers.Connect()

	// create fiber app
	app := fiber.New()

	// setup routes
	routes.Setup(routes.SetupOptions{Version: 1, App: app})

	app.Listen(":3000")
}
