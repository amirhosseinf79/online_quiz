package handler

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"github.com/amirhosseinf79/online_quiz/internal/service"
	"github.com/gofiber/fiber/v3"
)

type quizHandler struct {
	quizService service.QuizService
}

func NewQuizHandler(quizService service.QuizService) QuizHandler {
	return &quizHandler{
		quizService: quizService,
	}
}

func (h *quizHandler) GetAllQuizzes(c fiber.Ctx) error {
	var filter dto.QuizFilter
	response, err := dto.ValidateQueryParams(&filter, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	filter.PageFilter.SetDefaults()
	quizList, total, err := h.quizService.GetAllByFilter(filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": quizList,
		"total":  total,
	})
}

func (h *quizHandler) GetQuizById(c fiber.Ctx) error {
	quiz := c.Locals("quiz").(*models.Quiz)
	return c.Status(fiber.StatusOK).JSON(quiz)
}

func (h *quizHandler) CreateQuiz(c fiber.Ctx) error {
	var quiz dto.QuizCreate
	response, err := dto.ValidateRequestBody(&quiz, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	createdQuiz, err := h.quizService.CreateQuiz(quiz)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(createdQuiz)
}

func (h *quizHandler) UpdateQuiz(c fiber.Ctx) error {
	var quiz dto.QuizUpdate
	response, err := dto.ValidateRequestBody(&quiz, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	updatedQuiz, err := h.quizService.UpdateQuiz(quiz)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(updatedQuiz)
}

func (h *quizHandler) DeleteQuiz(c fiber.Ctx) error {
	quizID := fiber.Params[uint](c, "id")
	if quizID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Quiz ID is required",
		})
	}
	err := h.quizService.DeleteQuiz(quizID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
