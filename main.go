package main

import (
	"jwt/database"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
