package service

import (
	"errors"

	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"gorm.io/gorm"
)

type userAnswerService struct {
	userAnswerRepo repository.UserAnswerRepository
}

func NewUserAnswerService(userAnswerRepo repository.UserAnswerRepository) UserAnswerService {
	return &userAnswerService{
		userAnswerRepo: userAnswerRepo,
	}
}

func (s *userAnswerService) AddAnswer(fields dto.AddUserAnswer) error {
	answerM, err := s.userAnswerRepo.GetAnswer(fields.ResultID, fields.QuestionID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		answerM.AnswerID = fields.AnswerID
		return s.userAnswerRepo.Update(answerM)
	}

	if err != nil {
		return err
	}

	userAnswer := &models.UserAnswer{
		UserResultID: fields.ResultID,
		QuestionID:   fields.QuestionID,
		AnswerID:     fields.AnswerID,
	}

	return s.userAnswerRepo.Create(userAnswer)
}
