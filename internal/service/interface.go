package service

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
)

type UserService interface {
	RegisterUser(creds dto.UserRegister) (*models.User, error)
	LoginUser(creds dto.UserLogin) (*models.Token, error)
}

type TokenService interface {
	GenerateToken(userId uint) (*models.Token, error)
	RefreshToken(refreshToken string) (*models.Token, error)
}
