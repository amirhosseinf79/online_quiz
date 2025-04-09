package service

import (
	"time"

	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
)

type quizService struct {
	quizRepo repository.QuizRepository
}

func NewQuizService(quizRepo repository.QuizRepository) QuizService {
	return &quizService{
		quizRepo: quizRepo,
	}
}

func (s *quizService) GetAllByFilter(filter dto.QuizFilter) (quizzes []*models.Quiz, total int64, err error) {
	quizzes, total, err = s.quizRepo.GetAllByFilter(filter)
	return
}

func (s *quizService) GetQuizById(id uint) (quiz *models.Quiz, err error) {
	quiz, err = s.quizRepo.GetByID(id)
	return
}

func (s *quizService) CreateQuiz(quiz dto.QuizCreate) (*models.Quiz, error) {
	quizModel := &models.Quiz{
		Name: quiz.Name,
	}
	valid, err := s.quizRepo.CheckQuizDate(quiz.StartAt, quiz.EndAt)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, dto.ErrQuizDateNotValid

	}

	if err = s.quizRepo.Create(quizModel); err != nil {
		return nil, err
	}

	return quizModel, nil
}

func (s *quizService) UpdateQuiz(quiz dto.QuizUpdate) (*models.Quiz, error) {
	valid, err := s.quizRepo.CheckQuizDate(quiz.StartAt, quiz.EndAt)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, dto.ErrQuizDateNotValid
	}

	startAt, err := time.Parse(time.RFC3339, quiz.StartAt)
	if err != nil {
		return nil, err
	}

	endAt, err := time.Parse(time.RFC3339, quiz.EndAt)
	if err != nil {
		return nil, err
	}

	quizModel := &models.Quiz{
		ID:       quiz.ID,
		Name:     quiz.Name,
		Duration: quiz.Duration,
		StartAt:  startAt,
		EndAt:    endAt,
	}

	if err = s.quizRepo.Update(quizModel); err != nil {
		return nil, err
	}

	return quizModel, nil
}

func (s *quizService) DeleteQuiz(id uint) error {
	return s.quizRepo.Delete(id)
}
