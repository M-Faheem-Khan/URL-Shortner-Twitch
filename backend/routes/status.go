package routes

import "github.com/gofiber/fiber/v2"

func Status(c *fiber.Ctx) error {
	// Return
	// the status of the server
	// the connection status of db

	return c.SendString("Server is up and running!")
}

// EOF
