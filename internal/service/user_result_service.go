package service

import (
	"errors"

	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"gorm.io/gorm"
)

type userResultService struct {
	userResultRepo repository.UserResultRepository
	userAnswerRepo repository.UserAnswerRepository
	questionRepo   repository.QuestionRepository
}

func NewUserResultService(
	userResultRepo repository.UserResultRepository,
	userAnswerRepo repository.UserAnswerRepository,
	questionRepo repository.QuestionRepository,
) UserResultService {
	return &userResultService{
		userResultRepo: userResultRepo,
		userAnswerRepo: userAnswerRepo,
		questionRepo:   questionRepo,
	}
}

func (s *userResultService) GetByID(id uint) (*models.UserResult, error) {
	result, err := s.userResultRepo.GetByID(id)
	return result, err
}

func (s *userResultService) CreateOrGet(fields dto.ResultCreate) (*models.UserResult, error) {
	result, err := s.userResultRepo.GetByQuizUSerID(fields.QuizID, fields.UserID)
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

	err = s.userResultRepo.Create(newResult)
	if err != nil {
		return nil, err
	}
	result, err = s.userResultRepo.GetByQuizUSerID(fields.QuizID, fields.UserID)
	if err == nil {
		return result, nil
	}

	return newResult, nil
}

func (s *userResultService) UpdateResult(fields dto.ResultCreate) (result *models.UserResult, err error) {
	result, err = s.CreateOrGet(fields)
	if err != nil {
		return
	}
	totalQuestions, err := s.questionRepo.TotalQuestions(result.QuizID)
	if err != nil {
		return
	}
	trueAnswers, err := s.userAnswerRepo.GetCurrectAnsCount(result.ID)
	if err != nil {
		return
	}
	score := (float64(trueAnswers) / float64(totalQuestions)) * 100
	if score == result.Score {
		return
	}
	err = s.userResultRepo.UpdateScore(result.ID, score)
	return
}
