package handler

import (
	"gar/app/search"
	"github.com/gofiber/fiber/v3"
)

func Search(c fiber.Ctx) error {
	b := new(search.RequestBody)

	return c.JSON(b)
}
