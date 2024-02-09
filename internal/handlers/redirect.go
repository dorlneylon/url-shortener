package handlers

import (
	"github.com/gofiber/fiber/v2"
	"url-shortener/internal/models"
	"url-shortener/internal/storage"
)

func HandleRedirect(c *fiber.Ctx, mgo *storage.Mongo) error {
	alias := c.Params("alias")
	short, err := mgo.GetByShortened(alias)

	if err != nil {
		return c.JSON(fiber.Map{"err": err.Error()})
	}

	_, _ = mgo.IncrementClicks(alias)
	return c.Redirect(short.(models.Shortening).Url, 301)
}
