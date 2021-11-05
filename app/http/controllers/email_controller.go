package controllers

import (
	"latest/config/email"
	"latest/dto"
	"latest/infra/repository"
	"latest/pkg/logger"
	"latest/pkg/validate"
	"latest/usecases/dispatcher"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

type EmailController struct{}

func NewEmailController() Controller {
	return &EmailController{}
}

func (e *EmailController) MultipartDispatch(c *fiber.Ctx) error {

	var req dto.MultiPartEmailDTO

	files, err := c.MultipartForm()

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&dto.ErrorDTO{
			Code:    400,
			Message: err.Error(),
		})
	}

	defer func() {
		err := files.RemoveAll()
		if err != nil {
			logger.Instance().Warn(err.Error())
		}
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
			Message: err.Error(),
		})
	}

	for _, file := range files.File {
		for range file {
			req.Attachment = append(req.Attachment, file...)
		}
	}

	svc := dispatcher.NewUsecase(repository.NewMailRepository(email.GetInstance()))

	sendErr := svc.MultipartAttachments(&req)

	if sendErr != nil {
		return c.Status(sendErr.Code()).JSON(fiber.Map{"message": sendErr.Message()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Email was dispatched"})
}

func (e *EmailController) EncondedAttachments(c *fiber.Ctx) error {

	var req dto.EmailDTO

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err = validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusOK).JSON(&dto.ErrorDTO{
			Message: err.Error(),
		})
	}

	svc := dispatcher.NewUsecase(repository.NewMailRepository(email.GetInstance()))

	sendErr := svc.Base64Attachments(&req)

	if sendErr != nil {
		return c.Status(sendErr.Code()).JSON(fiber.Map{"message": sendErr.Message()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Email was dispatched"})
}
