package middleware

import (
	"github.com/amirhosseinf79/online_quiz/internal/service"
	"github.com/gofiber/fiber/v3"
)

type rollMiddleware struct {
	userService service.UserService
}

func NewRollMiddleware(userService service.UserService) RollMiddleware {
	return &rollMiddleware{userService: userService}
}

func (roll *rollMiddleware) AdminRequired(c fiber.Ctx) error {
	userM, err := roll.userService.GetUserById(c.Locals("userId").(uint))
	if err != nil {
		return err
	}
	if userM.Role != "admin" {
		return fiber.NewError(fiber.StatusForbidden, "Admin role required")
	}
	return c.Next()
}
