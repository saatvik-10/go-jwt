package main

import (
	"jwt/database"
	"jwt/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
