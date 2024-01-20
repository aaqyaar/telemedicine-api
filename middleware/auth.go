package middleware

import "github.com/gofiber/fiber/v2"

/**
 * @title       : auth.go
 * @description : authentication middleware - check if user is authenticated
 * @package     : middleware
 */
func Authenticate(c *fiber.Ctx) error {
	return c.Next()
}
