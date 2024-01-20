package routes

import (
	"github.com/aaqyaar/telemedicine/controller"
	"github.com/aaqyaar/telemedicine/middleware"
	"github.com/gofiber/fiber/v2"
)

/**
 * @title       : routes.go
 * @description : initialize routes for the application
 * @package     : routes
 */
func Setup(app *fiber.App) {

	app.Use(middleware.Authenticate)

	app.Get("/", controller.GetUsers)
	app.Get("/:id", controller.GetUser)

}
