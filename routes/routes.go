package routes

import (
	"import/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/stocks", controllers.Index)
	api.Get("/health-check", controllers.HealthCheck)
}
