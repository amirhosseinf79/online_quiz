package handler

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"github.com/amirhosseinf79/online_quiz/internal/service"
	"github.com/gofiber/fiber/v3"
)

type questionHandler struct {
	questionService service.QuestionService
}

func NewQuestionHandler(questionService service.QuestionService) QuestionHandler {
	return &questionHandler{
		questionService: questionService,
	}
}

func (questionHandler *questionHandler) GetAllQuestions(c fiber.Ctx) error {
	filter := dto.QuestionFilter{}
	response, err := dto.ValidateQueryParams(&filter, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	filter.QuizID = c.Locals("quiz").(*models.Quiz).ID
	filter.PageFilter.SetDefaults()
	questions, total, err := questionHandler.questionService.GetAllByFilter(filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "Failed to retrieve questions",
		})
	}

	return c.JSON(fiber.Map{
		"result": questions,
		"total":  total,
	})
}

func (questionHandler *questionHandler) GetQuestionById(c fiber.Ctx) error {
	id := fiber.Params[uint](c, "id")
	if id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Invalid question ID",
		})
	}

	question, err := questionHandler.questionService.GetQuestionById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Message: "Question not found",
		})
	}
	return c.JSON(question)
}

func (questionHandler *questionHandler) CreateQuestion(c fiber.Ctx) error {
	var question dto.QuestionCreate
	response, err := dto.ValidateRequestBody(&question, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	newQuestion, err := questionHandler.questionService.CreateQuestion(question)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(newQuestion)
}

func (questionHandler *questionHandler) UpdateQuestion(c fiber.Ctx) error {
	var question dto.QuestionUpdate
	response, err := dto.ValidateRequestBody(&question, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	updatedQuestion, err := questionHandler.questionService.UpdateQuestion(question)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(updatedQuestion)
}

func (questionHandler *questionHandler) DeleteQuestion(c fiber.Ctx) error {
	id := fiber.Params[uint](c, "id")
	if id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Invalid question ID",
		})
	}

	err := questionHandler.questionService.DeleteQuestion(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
