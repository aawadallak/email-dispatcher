package controllers

import (
	"latest/dto"
	"latest/pkg/validate"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

type EmailController struct{}

func NewEmailController() Controller {
	return &EmailController{}
}

func (e *EmailController) DispatchEmail(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pong"})
}

func (e *EmailController) MultiPartDispatch(c *fiber.Ctx) error {

	var dto dto.MultiPartEmailDTO

	files, err := c.MultipartForm()

	defer files.RemoveAll()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = mapstructure.Decode(files.Value, &dto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err = validate.Struct(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	for _, file := range files.File {
		for range file {
			dto.Attachment = append(dto.Attachment, file...)
		}
	}

	message, err := dto.Convert2Entity()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": message})
}
