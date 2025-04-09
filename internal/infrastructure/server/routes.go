package server

import "github.com/gofiber/fiber/v3"

func (s server) InitRoutes() {
	api := s.app.Group("/api/v1")

	s.initAuthRoutes(api)
	s.initQuizRoutes(api)
	s.initQuestionRoutes(api)
	s.initAnswerRoutes(api)
}

func (s server) initAuthRoutes(api fiber.Router) {
	api.Post("/auth/register", s.userHandler.RegisterUser)
	api.Post("/auth/login", s.userHandler.LoginUser)
}

func (s server) initQuizRoutes(api fiber.Router) {
	api.Get("/quiz", s.quizHandler.GetAllQuizzes)
	api.Use(s.tokenMiddleware.CheckTokenAuth)
	api.Get("/quiz/:id", s.quizHandler.GetQuizById, s.quizMiddleware.QuizDateValid)
	api.Get("/quiz/:id/questions", s.questionHandler.GetAllQuestions, s.quizMiddleware.QuizDateValid)
	api.Post("/quiz/add", s.quizHandler.CreateQuiz, s.rollMiddleware.AdminRequired)
	api.Put("/quiz/edit", s.quizHandler.UpdateQuiz, s.rollMiddleware.AdminRequired)
	api.Delete("/quiz/:id", s.quizHandler.DeleteQuiz, s.rollMiddleware.AdminRequired)
}

func (s server) initQuestionRoutes(api fiber.Router) {
	api.Get("/question/:id", s.questionHandler.GetQuestionById, s.rollMiddleware.AdminRequired)
	api.Post("/question/add", s.questionHandler.CreateQuestion, s.rollMiddleware.AdminRequired)
	api.Put("/question/edit", s.questionHandler.UpdateQuestion, s.rollMiddleware.AdminRequired)
	api.Delete("/question/:id", s.questionHandler.DeleteQuestion, s.rollMiddleware.AdminRequired)
}

func (s server) initAnswerRoutes(api fiber.Router) {
	api.Put("/answer/edit", s.answerHandler.UpdateAnswer, s.rollMiddleware.AdminRequired)
}
