package main

import (
	"github.com/amirhosseinf79/online_quiz/internal/application/handler"
	"github.com/amirhosseinf79/online_quiz/internal/application/middleware"
	"github.com/amirhosseinf79/online_quiz/internal/infrastructure/database"
	"github.com/amirhosseinf79/online_quiz/internal/infrastructure/persistence"
	"github.com/amirhosseinf79/online_quiz/internal/infrastructure/server"
	"github.com/amirhosseinf79/online_quiz/internal/service"
)

func main() {
	db := database.NewDB("host=localhost user=postgres password=Amir2001 dbname=online_quiz port=5432 sslmode=disable TimeZone=Asia/Tehran")
	// database.AddMockData(db)
	userRepo := persistence.NewUserRepository(db)
	quizRepo := persistence.NewQuizRepository(db)
	questionRepo := persistence.NewQuestionRepository(db)
	answerRepo := persistence.NewAnswerRepository(db)
	tokenRepo := persistence.NewTokenRepository(db)
	userResultRepo := persistence.NewUserResultRepo(db)
	userAnsRepo := persistence.NewUserAnswerRepo(db)

	tokenService := service.NewTokenService(tokenRepo)
	userService := service.NewUserService(userRepo, tokenService)
	quizService := service.NewQuizService(quizRepo)
	questionService := service.NewQuestionService(questionRepo)
	answerService := service.NewAnswerService(answerRepo)
	userResultService := service.NewUserResultService(userResultRepo)
	userAnsService := service.NewUserAnswerService(userAnsRepo)

	tokenMiddleware := middleware.NewAuthMiddleware(tokenService)
	rollMiddleware := middleware.NewRollMiddleware(userService)
	quizMiddleware := middleware.NewQuizMiddleware(userService, quizService)
	answerCheckMiddleware := middleware.NewAnswerMiddleware(userResultService)

	userHandler := handler.NewUserHandler(userService)
	quizHandler := handler.NewQuizHandler(quizService)
	questionHandler := handler.NewQuestionHandler(questionService)
	answerHandler := handler.NewAnswerHandler(answerService)
	tokenHandler := handler.NewTokenHandler(tokenService)
	userResultHandler := handler.NewUserResultHandler(userResultService)
	userAnswerHandler := handler.NewUserAnsHandler(userAnsService)

	server := server.NewServer(
		tokenMiddleware,
		quizMiddleware,
		rollMiddleware,
		answerCheckMiddleware,
		userHandler,
		quizHandler,
		questionHandler,
		answerHandler,
		tokenHandler,
		userResultHandler,
		userAnswerHandler,
	)
	server.InitServer()
	server.InitRoutes()
	server.Start()
}
