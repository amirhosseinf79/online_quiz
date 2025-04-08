package service

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
)

type answerService struct {
	answerRepo repository.AnswerRepository
}

func NewAnswerService(answerRepo repository.AnswerRepository) AnswerService {
	return &answerService{
		answerRepo: answerRepo,
	}
}

func (as *answerService) UpdateAnswer(fields dto.AnswerUpdate) (answer *models.Answer, err error) {
	answerM := &models.Answer{
		ID:        fields.ID,
		Text:      fields.Text,
		IsCorrect: fields.IsCorrect,
	}
	err = as.answerRepo.Update(answerM)
	answer = answerM
	return
}
