package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"url-shortener/internal/models"
	"url-shortener/internal/net/auth"
	"url-shortener/internal/storage"
)

func HandleDelete(c *fiber.Ctx, memcached *storage.Memcached, mgo *storage.Mongo) error {
	var payload models.UpdateRequest

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	sender, err := auth.DecodeJWT(payload.Jwt)
	if err != nil {
		return c.JSON(fiber.Map{"err": err.Error()})
	}

	author, err := mgo.GetByShortened(payload.Alias)

	if err != nil {
		return c.JSON(fiber.Map{"err": err.Error()})
	}

	if author.(models.Shortening).Author != sender {
		return c.JSON(fiber.Map{"err": errors.New("Not an author")})
	}

	_ = memcached.Delete(payload.Alias)
	_, err = mgo.DeleteShortening(payload.Alias)

	return c.JSON(fiber.Map{
		"err": err,
	})
}
