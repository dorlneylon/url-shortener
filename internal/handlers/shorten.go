package handlers

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"url-shortener/internal/models"
	"url-shortener/internal/net/auth"
	"url-shortener/internal/shorten"
	"url-shortener/internal/storage"
)

func HandleShorten(c *fiber.Ctx, mgo *storage.Mongo) error {
	var (
		payload models.ShortRequest
	)

	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(fiber.Map{"err": err.Error()})
	}

	author, err := auth.DecodeJWT(payload.Jwt)

	if err != nil {
		return c.JSON(fiber.Map{"err": err.Error()})
	}

	model := models.Shortening{
		Alias:     shorten.Shorten(payload.Url),
		Url:       payload.Url,
		Author:    author,
		Clicks:    0,
		Timestamp: time.Now().Unix(),
	}

	_, err = mgo.InsertShortening(model)

	if err != nil {
		return c.JSON(fiber.Map{"err": err.Error()})
	}

	return c.JSON(fiber.Map{
		"alias": model.Alias,
	})
}
