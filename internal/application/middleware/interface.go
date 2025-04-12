package middleware

import "github.com/gofiber/fiber/v3"

type TokenMiddleware interface {
	CheckTokenAuth(c fiber.Ctx) error
}

type RollMiddleware interface {
	AdminRequired(c fiber.Ctx) error
}

type QuizMiddleware interface {
	QuizDateValid(c fiber.Ctx) error
}

type AnswerMiddleware interface {
	CheckAnswer(c fiber.Ctx) error
}
