package routes

import (
	"fmt"

	"github.com/M-Faheem-Khan/Twitch-URL-Shortener/backend/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func Expand(c *fiber.Ctx) error {
	// get /:shortID
	// check if shortID is valid
	// if valid: redirect to original url
	// else: error

	shortID := c.Params("shortID")

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortID).Result()
	if err == redis.Nil {
		fmt.Println("Unable to get shortID ", shortID)
		fmt.Print(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "invalid shortId",
		})
	} else if err != nil {
		fmt.Println("Unable to get shortID ", shortID)
		fmt.Print(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.Redirect(val, 301)
}

// EOF
