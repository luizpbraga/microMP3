package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizpbraga/microMP3/services/gateway/src/controller"
)

func LoadRoutes(app *fiber.App) {
	app.Post("/login", controller.Login)
}
