package service

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
)

type UserService interface {
	RegisterUser(creds dto.UserRegister) (*models.Token, error)
	LoginUser(creds dto.UserLogin) (*models.Token, error)
}

type TokenService interface {
	GenerateToken(userId uint) (*models.Token, error)
	RefreshToken(refreshToken string) (*models.Token, error)
}

type QuizService interface {
	GetAllByFilter(filter dto.QuizFilter) ([]*models.Quiz, int64, error)
	GetQuizById(id uint) (*models.Quiz, error)
	CreateQuiz(quiz dto.QuizCreate) (*models.Quiz, error)
}

type QuestionService interface {
	GetAllByFilter(filter dto.QuestionFilter) ([]*models.Question, int64, error)
	GetQuestionById(id uint) (*models.Question, error)
	CreateQuestion(question dto.QuestionCreate) (*models.Question, error)
}

type AnswerService interface {
	UpdateAnswer(fields dto.AnswerUpdate) (*models.Answer, error)
}
