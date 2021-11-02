package routes

import (
	"latest/app/http/controllers"

	"github.com/gofiber/fiber/v2"
)

// Return an setup router for fiber application
func SetRoute(router *fiber.App) *fiber.App {

	controller := controllers.NewEmailController()

	// Define the group router which will be used on project
	api := router.Group("api/v1")

	// Router responsable to dispatch emails, recive a request body with content
	api.Post("/dispatch", controller.DispatchEmail)

	// Health check route, use for check application is running
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "running..."})
	})

	return router
}
