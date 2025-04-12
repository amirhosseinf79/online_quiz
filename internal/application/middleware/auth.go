package middleware

import (
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"github.com/amirhosseinf79/online_quiz/internal/service"
	"github.com/gofiber/fiber/v3"
)

type authMiddleware struct {
	tokenService service.TokenService
}

func NewAuthMiddleware(tokenService service.TokenService) TokenMiddleware {
	return &authMiddleware{tokenService: tokenService}
}

func (a *authMiddleware) CheckTokenAuth(c fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{Message: "No token provided"})
	}
	tokenM, err := a.tokenService.GetByToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{Message: "Invalid token"})
	}
	c.Locals("userId", tokenM.UserID)
	return c.Next()
}
