package middleware

import (
	"fmt"
	"time"

	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"github.com/amirhosseinf79/online_quiz/internal/service"
	"github.com/gofiber/fiber/v3"
)

type quizMiddleware struct {
	userService service.UserService
	quizService service.QuizService
}

func NewQuizMiddleware(userService service.UserService, quizService service.QuizService) QuizMiddleware {
	return &quizMiddleware{
		userService: userService,
		quizService: quizService,
	}
}

func (quiz *quizMiddleware) QuizDateValid(c fiber.Ctx) error {
	quizID := fiber.Params[uint](c, "id")
	if quizID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Quiz ID is required",
		})
	}

	quizM, err := quiz.quizService.GetQuizById(quizID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Message: "Quiz not found",
		})
	}
	fmt.Println(quizM)
	c.Locals("quiz", quizM)
	user, err := quiz.userService.GetUserById(c.Locals("userId").(uint))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
			Message: "User not found",
		})
	}

	if user.Role != "admin" && quizM.StartAt.After(time.Now()) {
		return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
			Message: dto.ErrQuizNotStarted.Error(),
		})
	}

	if user.Role != "admin" && quizM.EndAt.Before(time.Now()) {
		return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
			Message: dto.ErrQuizEnded.Error(),
		})
	}

	return c.Next()
}
