package handler

import (
	"gar/app/search"
	"github.com/gofiber/fiber/v3"
)

func Categories(c fiber.Ctx) error {
	return c.JSON(search.Categories)
}
