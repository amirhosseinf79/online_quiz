package repository

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
)

type QuestionRepository interface {
	GetAllByFilter(filter dto.QuestionFilter) ([]*models.Question, int64, error)
	TotalQuestions(quizID uint) (int64, error)
	GetByID(id uint) (*models.Question, error)
	Create(quiz *models.Question) error
	Update(quiz *models.Question) error
	Delete(id uint) error
}
