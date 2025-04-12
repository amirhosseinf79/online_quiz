package handler

import "github.com/gofiber/fiber/v3"

type UserHandler interface {
	RegisterUser(c fiber.Ctx) error
	LoginUser(c fiber.Ctx) error
}

type TokenHandler interface {
	RefreshToken(c fiber.Ctx) error
}

type QuizHandler interface {
	GetAllQuizzes(c fiber.Ctx) error
	GetQuizById(c fiber.Ctx) error
	CreateQuiz(c fiber.Ctx) error
	UpdateQuiz(c fiber.Ctx) error
	DeleteQuiz(c fiber.Ctx) error
}

type QuestionHandler interface {
	GetAllQuestions(c fiber.Ctx) error
	GetQuestionById(c fiber.Ctx) error
	CreateQuestion(c fiber.Ctx) error
	UpdateQuestion(c fiber.Ctx) error
	DeleteQuestion(c fiber.Ctx) error
}

type AnswerHandler interface {
	UpdateAnswer(c fiber.Ctx) error
}

type UserResultHandler interface {
	GetQuizResultDetails(c fiber.Ctx) error
}

type UserAnsHandler interface {
	AddAnswer(c fiber.Ctx) error
}
