package service

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"github.com/amirhosseinf79/online_quiz/pkg"
)

type userService struct {
	userRepo     repository.UserRepository
	tokenService TokenService
}

func NewUserService(userRepo repository.UserRepository, tokenService TokenService) UserService {
	return &userService{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (s *userService) RegisterUser(creds dto.UserRegister) (*models.Token, error) {
	exists, err := s.userRepo.CheckEmailExists(creds.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, dto.ErrEmailExists
	}

	hashedPassword, err := pkg.HashPassword(creds.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email:     creds.Email,
		Password:  hashedPassword,
		FirstName: creds.FirstName,
		LastName:  creds.LastName,
	}

	err = s.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	token, err := s.tokenService.GenerateToken(user.ID)
	return token, err
}

func (s *userService) LoginUser(creds dto.UserLogin) (*models.Token, error) {
	user, err := s.userRepo.GetByEmail(creds.Email)
	if err != nil {
		return nil, dto.ErrInvalidCredentials
	}

	if valid := user.ValidatePassword(creds.Password); !valid {
		return nil, dto.ErrInvalidCredentials
	}

	token, err := s.tokenService.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return token, nil
}
