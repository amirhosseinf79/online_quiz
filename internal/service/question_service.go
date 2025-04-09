package service

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
)

type questionService struct {
	questionRepo repository.QuestionRepository
}

func NewQuestionService(questionRepo repository.QuestionRepository) QuestionService {
	return &questionService{
		questionRepo: questionRepo,
	}
}

func (qs *questionService) GetAllByFilter(filter dto.QuestionFilter) (questions []*models.Question, total int64, err error) {
	questions, total, err = qs.questionRepo.GetAllByFilter(filter)
	return
}

func (qs *questionService) GetQuestionById(id uint) (question *models.Question, err error) {
	question, err = qs.questionRepo.GetByID(id)
	return
}

func (qs *questionService) CreateQuestion(question dto.QuestionCreate) (createdQuestion *models.Question, err error) {
	questionM := models.Question{
		QuizID:   question.QuizID,
		Question: question.Text,
	}
	if len(question.Answers) < 4 {
		err = dto.ErrNoAnswersProvided
		return
	}

	var correctAnswersCount int
	for _, answer := range question.Answers {
		questionM.Answers = append(questionM.Answers, models.Answer{Text: answer.Text, IsCorrect: answer.IsCorrect})
		if answer.IsCorrect {
			correctAnswersCount++
		}
	}

	if correctAnswersCount == 1 {
		err = qs.questionRepo.Create(&questionM)
		createdQuestion = &questionM
	} else {
		err = dto.ErrMultipleCorrectAnswers
	}
	return
}

func (qs *questionService) UpdateQuestion(question dto.QuestionUpdate) (questionM *models.Question, err error) {
	questionM = &models.Question{
		ID:       question.ID,
		QuizID:   question.QuizID,
		Question: question.Text,
	}
	if len(question.Answers) < 4 {
		err = dto.ErrNoAnswersProvided
		return
	}

	var correctAnswersCount int
	for _, answer := range question.Answers {
		questionM.Answers = append(questionM.Answers, models.Answer{Text: answer.Text, IsCorrect: answer.IsCorrect})
		if answer.IsCorrect {
			correctAnswersCount++
		}
	}

	if correctAnswersCount == 1 {
		err = qs.questionRepo.Update(questionM)
	} else {
		err = dto.ErrMultipleCorrectAnswers
	}
	return
}

func (qs *questionService) DeleteQuestion(id uint) (err error) {
	err = qs.questionRepo.Delete(id)
	return
}
