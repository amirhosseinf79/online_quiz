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
	answerM, err := as.answerRepo.GetByID(fields.ID)
	if err != nil {
		return nil, err
	}

	answerM.Text = fields.Text
	answerM.IsCorrect = fields.IsCorrect

	if fields.IsCorrect {
		// Set other answers for the same question to incorrect
		err = as.answerRepo.SetOtherAnswersIncorrect(answerM.QuestionID, answerM.ID)
		if err != nil {
			return nil, err
		}
	} else {
		// Ensure at least one correct answer exists
		count, err := as.answerRepo.CountCorrectAnswers(answerM.QuestionID, answerM.ID)
		if err != nil {
			return nil, err
		}
		if count == 0 {
			return nil, dto.ErrMultipleCorrectAnswers
		}
	}

	err = as.answerRepo.Update(answerM)
	answer = answerM
	return
}
