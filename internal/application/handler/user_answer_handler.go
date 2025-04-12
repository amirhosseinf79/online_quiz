package handler

import (
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"github.com/amirhosseinf79/online_quiz/internal/service"
	"github.com/gofiber/fiber/v3"
)

type userAnsHandler struct {
	userAnsService service.UserAnswerService
}

func NewUserAnsHandler(userAnsService service.UserAnswerService) UserAnsHandler {
	return &userAnsHandler{userAnsService: userAnsService}
}

func (h *userAnsHandler) AddAnswer(c fiber.Ctx) error {
	var fields dto.AddUserAnswer
	response, err := dto.ValidateRequestBody(&fields, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err = h.userAnsService.AddAnswer(fields)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(dto.ErrorResponse{
		Message: "ok",
	})
}
