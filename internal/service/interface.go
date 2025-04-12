package service

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
)

type UserService interface {
	RegisterUser(creds dto.UserRegister) (*models.Token, error)
	LoginUser(creds dto.UserLogin) (*models.Token, error)
	GetUserById(id uint) (*models.User, error)
}

type TokenService interface {
	GenerateToken(userId uint) (*models.Token, error)
	RefreshToken(refreshToken string) (*models.Token, error)
	GetByToken(token string) (*models.Token, error)
}

type QuizService interface {
	GetAllByFilter(filter dto.QuizFilter) ([]*models.Quiz, int64, error)
	GetQuizById(id uint) (*models.Quiz, error)
	CreateQuiz(quiz dto.QuizCreate) (*models.Quiz, error)
	UpdateQuiz(quiz dto.QuizUpdate) (*models.Quiz, error)
	DeleteQuiz(id uint) error
}

type QuestionService interface {
	GetAllByFilter(filter dto.QuestionFilter) ([]*models.Question, int64, error)
	GetQuestionById(id uint) (*models.Question, error)
	CreateQuestion(question dto.QuestionCreate) (*models.Question, error)
	UpdateQuestion(question dto.QuestionUpdate) (*models.Question, error)
	DeleteQuestion(id uint) error
}

type AnswerService interface {
	UpdateAnswer(fields dto.AnswerUpdate) (*models.Answer, error)
}

type UserResultService interface {
	CreateOrGet(fields dto.ResultCreate) (*models.UserResult, error)
	GetByID(id uint) (*models.UserResult, error)
}

type UserAnswerService interface {
	AddAnswer(fields dto.AddUserAnswer) error
}
