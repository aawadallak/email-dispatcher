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

	var req dto.MultiPartEmailDTO

	files, err := c.MultipartForm()

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&dto.ErrorDTO{
			Code:    400,
			Message: err.Error(),
		})
	}

	defer func() error {
		err := files.RemoveAll()
		if err != nil {
			return c.Status(fiber.StatusOK).JSON(&dto.ErrorDTO{
				Code:    400,
				Message: err.Error(),
			})
		}
		return nil
	}()

	err = mapstructure.Decode(files.Value, &req)

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&dto.ErrorDTO{
			Code:    400,
			Message: err.Error(),
		})
	}

	if err = validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusOK).JSON(&dto.ErrorDTO{
			Code:    400,
			Message: err.Error(),
		})
	}

	for _, file := range files.File {
		for range file {
			req.Attachment = append(req.Attachment, file...)
		}
	}

	message, errEntity := req.Convert2Entity()

	if errEntity != nil {
		return c.Status(fiber.StatusOK).JSON(errEntity)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": message})
}
