package routes

import (
	"fmt"
	"strings"
	"time"

	"github.com/M-Faheem-Khan/Twitch-URL-Shortener/backend/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type request struct {
	URL string `json:"url"`
}

type response struct {
	SHORTID string `json:"shortID"`
}

func generateShortID() string {
	shortID := strings.ReplaceAll(uuid.New().String(), "-", "")
	return shortID[:10]
}

func Shorten(c *fiber.Ctx) error {
	// [x] Given a url
	// [x] Generate a shortID for the url
	// [ ] Expire a url 24 hr or 7 days

	// Parsing Body
	body := new(request)
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "malformed body",
		})
	}

	// Checking if the url field is provided and is valid
	if body.URL == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "url field is not valid.",
		})
	}

	shortID := generateShortID()

	r := database.CreateClient(0)
	defer r.Close()

	val, _ := r.Get(database.Ctx, shortID).Result()
	if val != "" {
		fmt.Println("ShortID: ", shortID, " - already in use!")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	err = r.Set(database.Ctx, shortID, body.URL, 24*3600*time.Second).Err()
	if err != nil {
		fmt.Println("Error saving shortID ", shortID, " to redis.")
		fmt.Println(err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"url":     body.URL,
		"shortID": shortID,
	})
}

// EOF
