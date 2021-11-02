package controllers

import "github.com/gofiber/fiber/v2"

type EmailController struct{}

func NewEmailController() Controller {
	return &EmailController{}
}

func (e *EmailController) DispatchEmail(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pong"})
}
