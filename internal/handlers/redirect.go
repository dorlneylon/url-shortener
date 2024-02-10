package handlers

import (
	"github.com/gofiber/fiber/v2"
	"url-shortener/internal/models"
	"url-shortener/internal/storage"
)

func HandleRedirect(c *fiber.Ctx, memcached *storage.Memcached, mgo *storage.Mongo) error {
	alias := c.Params("alias")
	long, err := memcached.Get(alias)

	if err == nil {
		go mgo.IncrementClicks(alias)
		return c.Redirect(long, 301)
	}

	short, err := mgo.GetByShortened(alias)

	if err != nil {
		return c.JSON(fiber.Map{"err": err.Error()})
	}

	go mgo.IncrementClicks(alias)
	return c.Redirect(short.(models.Shortening).Url, 301)
}
