package utils

import "github.com/gofiber/fiber/v2"

type ApiResponse struct {
	StatusCode int
	Data       interface{}
}

func Success(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func Error(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"error":   message,
	})
}
