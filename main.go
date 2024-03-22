package main

import (
	"import/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := fiber.New()

	routes.SetupRoutes(app)

	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
