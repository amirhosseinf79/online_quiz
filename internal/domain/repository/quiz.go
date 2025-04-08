package repository

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
)

type QuizRepository interface {
	GetByID(id uint) (*models.Quiz, error)
	GetAllByFilter(filter dto.QuizFilter) ([]*models.Quiz, int64, error)
	Create(quiz *models.Quiz) error
	Update(quiz *models.Quiz) error
	Delete(id uint) error
}
