package main

import (
	"log"

	"poke-go/config" // Adjust this path based on your project structure
	"poke-go/routes" // Import the routes package

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Load environment variables
	cfg := config.LoadEnv()

	// Initialize Fiber app
	app := fiber.New()
	// Setup CORS middleware based on environment
	if cfg.Environment == "production" {
		// Production CORS configuration
		app.Use(cors.New(cors.Config{
			AllowOrigins: "http://localhost, http://frontend", // Replace with your production frontend URL
			AllowHeaders: "Origin, Content-Type, Accept",
			AllowMethods: "GET, POST, PATCH, PUT, DELETE, OPTIONS",
		}))
	} else {
		// Development CORS configuration
		app.Use(cors.New(cors.Config{
			AllowOrigins: "http://localhost:5173", // Allow requests from your frontend URL
			AllowHeaders: "Origin, Content-Type, Accept",
			AllowMethods: "GET, POST, PATCH, PUT, DELETE, OPTIONS",
		}))
	}

	// Setup routes
	routes.SetupRoutes(app, cfg)

	// Start the server and log the result
	log.Printf("Starting server on port %s...\n", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
