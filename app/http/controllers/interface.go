package controllers

import "github.com/gofiber/fiber/v2"

type Controller interface {
	MultipartDispatch(*fiber.Ctx) error
	EncondedAttachments(*fiber.Ctx) error
}
