package routes

import (
	"latest/app/http/controllers"

	"github.com/gofiber/fiber/v2"
)

// Return an setup router for fiber application
func SetRoute(router *fiber.App) *fiber.App {

	controller := controllers.NewEmailController()

	// Set main route group
	api := router.Group("api/v1")

	// Dispatch an email with attachments with JSON Body containing enconded base64 attachments
	api.Post("/dispatch", controller.EncondedAttachments)

	// Dispatch an email with attachments using multipart
	api.Post("/dispatch/attachment", controller.MultipartDispatch)

	// Health check route, use for check application is running
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "service is running"})
	})

	return router
}
