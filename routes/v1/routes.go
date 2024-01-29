package routes

import (
	"fmt"

	"github.com/aaqyaar/telemedicine/controller"
	"github.com/aaqyaar/telemedicine/middleware"
	"github.com/gofiber/fiber/v2"
)

type SetupOptions struct {
	Version int
	App     *fiber.App
}

/**
 * @title       : routes.go
 * @description : initialize routes for the application
 * @package     : routes
 */
func Setup(opts SetupOptions) {
	app := opts.App

	if opts.App == nil {
		panic("Fiber app is missing")
	}

	app.Use(middleware.Authenticate)

	version := fmt.Sprintf("/v%d", opts.Version)

	v1 := app.Group(version, func(c *fiber.Ctx) error {
		c.Set("X-VERSION", version)
		return c.Next()
	})

	v1.Get("/users", controller.GetUsers)
	v1.Post("/users", controller.CreateUser)

}
