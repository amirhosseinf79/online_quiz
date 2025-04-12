package server

import (
	"log"

	"github.com/amirhosseinf79/online_quiz/internal/application/handler"
	"github.com/amirhosseinf79/online_quiz/internal/application/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type server struct {
	app                *fiber.App
	tokenMiddleware    middleware.TokenMiddleware
	quizMiddleware     middleware.QuizMiddleware
	rollMiddleware     middleware.RollMiddleware
	ansCheckMiddleware middleware.AnswerMiddleware
	userHandler        handler.UserHandler
	quizHandler        handler.QuizHandler
	questionHandler    handler.QuestionHandler
	answerHandler      handler.AnswerHandler
	tokenHandler       handler.TokenHandler
	userResultHandler  handler.UserResultHandler
	userAnswerHandler  handler.UserAnsHandler
}

func NewServer(
	tokenMiddleware middleware.TokenMiddleware,
	quizMiddleware middleware.QuizMiddleware,
	rollMiddleware middleware.RollMiddleware,
	ansCheckMiddleware middleware.AnswerMiddleware,
	userHandler handler.UserHandler,
	quizHandler handler.QuizHandler,
	questionHandler handler.QuestionHandler,
	answerHandler handler.AnswerHandler,
	tokenHandler handler.TokenHandler,
	userResultHandler handler.UserResultHandler,
	userAnswerHandler handler.UserAnsHandler,
) *server {
	return &server{
		tokenMiddleware:    tokenMiddleware,
		quizMiddleware:     quizMiddleware,
		rollMiddleware:     rollMiddleware,
		ansCheckMiddleware: ansCheckMiddleware,
		userHandler:        userHandler,
		quizHandler:        quizHandler,
		questionHandler:    questionHandler,
		answerHandler:      answerHandler,
		tokenHandler:       tokenHandler,
		userResultHandler:  userResultHandler,
		userAnswerHandler:  userAnswerHandler,
	}
}

func (s *server) InitServer() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{}))
	s.app = app
}

func (s *server) Start() {
	err := s.app.Listen(":3000")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
