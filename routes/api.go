package routes

import (
	"app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Example route
	api.Get("/hello", controllers.HelloHandler)
}
