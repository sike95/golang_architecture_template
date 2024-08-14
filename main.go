package main

import (
	"golang_architectire_template/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		Prefork:       false, // Disable prefork
		CaseSensitive: true,  // Enable case-sensitive routing
		StrictRouting: true,  // Enable strict routing
	})

	// Set up routes
	router := app.Group("/api")
	routes.CreateRoutes(router)

	// Start the server on port 3000
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
