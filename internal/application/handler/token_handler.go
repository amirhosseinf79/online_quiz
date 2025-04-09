package handler

import (
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"github.com/amirhosseinf79/online_quiz/internal/service"
	"github.com/gofiber/fiber/v3"
)

type tokenHandler struct {
	tokenService service.TokenService
}

func NewTokenHandler(tokenService service.TokenService) TokenHandler {
	return &tokenHandler{
		tokenService: tokenService,
	}
}

func (h *tokenHandler) RefreshToken(c fiber.Ctx) error {
	var body dto.RefreshToken
	response, err := dto.ValidateRequestBody(&body, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	token, err := h.tokenService.RefreshToken(body.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(token)
}
