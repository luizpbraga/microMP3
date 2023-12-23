package controller

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	// JWT auth token
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "StatusUnauthorized access",
		})
	}
	return c.SendString("Hello, World 👋!")
}
