package repository

import "github.com/amirhosseinf79/online_quiz/internal/domain/models"

type TokenRepository interface {
	GetByToken(token string) (*models.Token, error)
	Create(token *models.Token) error
	Update(token *models.Token) error
	Delete(token string) error
}
