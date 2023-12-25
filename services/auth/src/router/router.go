package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizpbraga/microMP3/services/auth/src/controller"
)

func LoadRoutes(app *fiber.App) {
	app.Post("/login", controller.Login)
	app.Post("/validate", controller.Validate)
}
