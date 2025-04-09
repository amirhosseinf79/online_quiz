package repository

import "github.com/amirhosseinf79/online_quiz/internal/domain/models"

type AnswerRepository interface {
	GetByID(id uint) (*models.Answer, error)
	Update(answer *models.Answer) error
}
