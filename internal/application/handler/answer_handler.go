package handler

import (
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"github.com/amirhosseinf79/online_quiz/internal/service"
	"github.com/gofiber/fiber/v3"
)

type ansHandler struct {
	ansService service.AnswerService
}

func NewAnswerHandler(ansService service.AnswerService) AnswerHandler {
	return &ansHandler{
		ansService: ansService,
	}
}

func (ansHandler *ansHandler) UpdateAnswer(c fiber.Ctx) error {
	var fields dto.AnswerUpdate
	response, err := dto.ValidateRequestBody(&fields, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	answer, err := ansHandler.ansService.UpdateAnswer(fields)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(answer)
}
