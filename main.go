package main

import (
	"log"

	"app/config"
	"app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize database
	config.InitDatabase()

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
