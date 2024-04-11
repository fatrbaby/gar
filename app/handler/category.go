package handler

import (
	"gar/app/ent"
	"github.com/gofiber/fiber/v3"
)

func Categories(c fiber.Ctx) error {
	return c.JSON(ent.Categories)
}
