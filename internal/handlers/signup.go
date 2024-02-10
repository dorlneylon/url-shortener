package handlers

import (
	"url-shortener/internal/models"
	"url-shortener/internal/net/auth"
	"url-shortener/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func HandleSignUp(c *fiber.Ctx, mgo *storage.Mongo) error {
	var (
		payload models.SignRequest
	)

	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(fiber.Map{"err": err.Error()})
	}

	_, err := mgo.InsertUser(models.User{
		Name:     payload.Name,
		Password: payload.Password,
	})

	if err != nil {
		return c.JSON(fiber.Map{"err": err.Error()})
	}

	jwt, err := auth.MakeJWT(payload.Name)

	return c.JSON(fiber.Map{
		"jwt": jwt,
		"err": err,
	})
}
