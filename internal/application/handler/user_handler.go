package handler

import (
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"github.com/amirhosseinf79/online_quiz/internal/service"
	"github.com/gofiber/fiber/v3"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) RegisterUser(c fiber.Ctx) error {
	var body dto.UserRegister
	response, err := dto.ValidateRequestBody(&body, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	token, err := h.userService.RegisterUser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(token)
}

func (h *userHandler) LoginUser(c fiber.Ctx) error {
	var body dto.UserLogin
	response, err := dto.ValidateRequestBody(&body, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	token, err := h.userService.LoginUser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(token)
}
