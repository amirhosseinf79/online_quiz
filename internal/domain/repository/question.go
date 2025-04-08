package repository

import "github.com/amirhosseinf79/online_quiz/internal/domain/models"

type QuestionRepository interface {
	GetByID(id uint) (*models.Question, error)
	Create(quiz *models.Question) error
	Update(quiz *models.Question) error
	Delete(id uint) error
}
