package handler

import (
	"gar/app/ent"
	"github.com/gofiber/fiber/v2"
)

func Categories(c *fiber.Ctx) error {
	return c.JSON(ent.Categories)
}
