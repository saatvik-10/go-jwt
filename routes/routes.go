package routes

import (
	"github.com/gofiber/fiber/v3"
	"jwt/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
