package repository

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
)

type UserResultRepository interface {
	Create(result *models.UserResult) error
	Update(result *models.UserResult) error
	GetByID(id uint) (*models.UserResult, error)
	GetByQuizUSerID(quizID, userID uint) (*models.UserResult, error)
	GetAll(filter dto.ResultFilter) ([]*models.UserResult, int64, error)
	ResultExists(quizID, userID uint) (bool, error)
}
