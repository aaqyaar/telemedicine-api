package controller

import (
	"fmt"

	"github.com/aaqyaar/telemedicine/models"
	"github.com/aaqyaar/telemedicine/providers"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {

	var user models.User

	fmt.Println(user.FirstName)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	providers.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message": "User Created success",
	})
}

func GetUsers(c *fiber.Ctx) error {

	var users []models.User

	providers.DB.Find(&users)

	return c.JSON(fiber.Map{"data": users})
}

func GetUser(c *fiber.Ctx) error {
	return c.SendString("Single User")
}
