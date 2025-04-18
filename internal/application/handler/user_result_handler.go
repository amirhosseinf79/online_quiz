package handler

import (
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"github.com/amirhosseinf79/online_quiz/internal/service"
	"github.com/gofiber/fiber/v3"
)

type userResultHandler struct {
	userResultService service.UserResultService
}

func NewUserResultHandler(userResultService service.UserResultService) UserResultHandler {
	return &userResultHandler{userResultService: userResultService}
}

func (h *userResultHandler) GetQuizResultDetails(c fiber.Ctx) error {
	var filter dto.ResultCreate
	response, err := dto.ValidateQueryParams(&filter, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	filter.UserID = c.Locals("userId").(uint)

	userResult, err := h.userResultService.UpdateResult(filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "Faild to obtain result id",
		})
	}

	return c.JSON(userResult)
}
