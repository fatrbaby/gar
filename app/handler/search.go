package handler

import (
	"gar/app/ent"
	"github.com/gofiber/fiber/v2"
)

func Search(c *fiber.Ctx) error {
	b := new(ent.RequestBody)

	return c.JSON(b)
}
