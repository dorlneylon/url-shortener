package handlers

import (
	"github.com/gofiber/fiber/v2"
	"url-shortener/internal/models"
	"url-shortener/internal/net/auth"
	"url-shortener/internal/storage"
)

func HandleSignIn(c *fiber.Ctx, mgo *storage.Mongo) error {
	var (
		payload models.SignRequest
	)

	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(fiber.Map{"err": err.Error()})
	}

	user, err := mgo.GetUser(payload.Name)

	if err != nil {
		return c.JSON(fiber.Map{"err": err.Error()})
	}

	if user.(models.User).Password != payload.Password {
		return fiber.ErrUnauthorized
	}

	jwt, err := auth.MakeJWT(payload.Name)

	return c.JSON(fiber.Map{
		"jwt": jwt,
		"err": err,
	})
}
