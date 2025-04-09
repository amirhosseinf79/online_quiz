package service

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"github.com/amirhosseinf79/online_quiz/pkg"
)

type tokenService struct {
	repo repository.TokenRepository
}

func NewTokenService(repo repository.TokenRepository) TokenService {
	return &tokenService{repo: repo}
}

func (t *tokenService) GenerateToken(userId uint) (*models.Token, error) {
	token := &models.Token{
		UserID: userId,
	}

	err := t.repo.Create(token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (t *tokenService) RefreshToken(refreshToken string) (*models.Token, error) {
	tokenM, err := t.repo.GetByRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	tokenM.RefreshToken = pkg.GenerateToken()
	err = t.repo.Update(tokenM)
	return tokenM, err
}

func (t *tokenService) GetByToken(token string) (*models.Token, error) {
	tokenM, err := t.repo.GetByToken(token)
	if err != nil {
		return nil, err
	}
	return tokenM, nil
}
