package service

import (
	"errors"

	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"gorm.io/gorm"
)

type userResultService struct {
	repo repository.UserResultRepository
}

func NewUserResultService(repo repository.UserResultRepository) UserResultService {
	return &userResultService{repo: repo}
}

func (s *userResultService) GetByID(id uint) (*models.UserResult, error) {
	result, err := s.repo.GetByID(id)
	return result, err
}

func (s *userResultService) CreateOrGet(fields dto.ResultCreate) (*models.UserResult, error) {
	result, err := s.repo.GetByQuizUSerID(fields.QuizID, fields.UserID)
	if err == nil {
		return result, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	newResult := &models.UserResult{
		QuizID: fields.QuizID,
		UserID: fields.UserID,
	}

	err = s.repo.Create(newResult)
	if err != nil {
		return nil, err
	}
	result, err = s.repo.GetByQuizUSerID(fields.QuizID, fields.UserID)
	if err == nil {
		return result, nil
	}

	return newResult, nil
}
