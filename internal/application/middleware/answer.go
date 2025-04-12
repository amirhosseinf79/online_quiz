package middleware

import (
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"github.com/amirhosseinf79/online_quiz/internal/service"
	"github.com/gofiber/fiber/v3"
)

type answerMiddleware struct {
	userResultService service.UserResultService
}

func NewAnswerMiddleware(userResultService service.UserResultService) AnswerMiddleware {
	return &answerMiddleware{
		userResultService: userResultService,
	}
}

func (a *answerMiddleware) CheckAnswer(c fiber.Ctx) error {
	var filter dto.GetResult
	response, err := dto.ValidateRequestBody(&filter, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	result, err := a.userResultService.GetByID(filter.ResultID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	if result.UserID != c.Locals("userId").(uint) {
		return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
			Message: dto.ErrPermission.Error(),
		})
	}
	return c.Next()
}
