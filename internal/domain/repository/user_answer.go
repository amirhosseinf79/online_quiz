package repository

import "github.com/amirhosseinf79/online_quiz/internal/domain/models"

type UserAnswerRepository interface {
	Create(answer *models.UserAnswer) error
	Update(answer *models.UserAnswer) error
	GetAnswer(resultID, questionID uint) (*models.UserAnswer, error)
}
