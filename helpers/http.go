package helpers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Error(c *fiber.Ctx, code int, err error) error {
	return c.
		Status(code).
		JSON(fiber.Map{
			"error": err.Error(),
			"data":  nil,
		})
}

func Success(c *fiber.Ctx, data any) error {
	return c.
		Status(http.StatusOK).
		JSON(fiber.Map{
			"error": nil,
			"data":  data,
		})
}
