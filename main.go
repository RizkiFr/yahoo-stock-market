package main

import (
	"go-whatsapp/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
