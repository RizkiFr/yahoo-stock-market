package main

import (
	"go-whatsapp/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not specified
	}
	app := fiber.New()

	routes.SetupRoutes(app)

	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
