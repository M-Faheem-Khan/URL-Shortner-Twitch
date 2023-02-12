package main

import (
	"log"
	"os"

	"github.com/M-Faheem-Khan/Twitch-URL-Shortener/backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:shortID", routes.Expand)
	app.Get("/api/v1/status", routes.Status)
	app.Post("/api/v1/shorten", routes.Shorten)
}

func main() {
	// Loading Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup Web Server (Fiber)
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	setupRoutes(app)

	app.Listen(os.Getenv("SERVER_PORT"))
}

// EOF
