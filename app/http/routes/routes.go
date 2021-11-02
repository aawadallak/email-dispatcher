package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoute(router *fiber.App) *fiber.App {

	router.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pong"})
	})

	return router
}
