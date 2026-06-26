package utils

import (
	"github.com/gofiber/fiber/v2"
)

func ParseBody(c *fiber.Ctx, req interface{}) error {
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON body")
	}
	return nil
}

func ParseQuery(c *fiber.Ctx, req interface{}) error {

	if err := c.QueryParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	return nil
}
