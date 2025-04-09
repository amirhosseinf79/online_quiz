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
		Name:     quiz.Name,
		Duration: quiz.Duration,
	}

	valid, err := s.quizRepo.CheckQuizDate(quiz.StartAt, quiz.EndAt)
	if err != nil {
		return nil, err
	}

	if quiz.StartAt != "" {
		start_at, err := time.Parse(time.RFC3339, quiz.StartAt)
		if err != nil {
			return nil, err
		}
		quizModel.StartAt = start_at
	}

	if quiz.EndAt != "" {
		endAt, err := time.Parse(time.RFC3339, quiz.EndAt)
		if err != nil {
			return nil, err
		}
		quizModel.EndAt = endAt
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
	quizM, err := s.quizRepo.GetByID(quiz.ID)
	if err != nil {
		return nil, err
	}

	if quiz.Name != "" {
		quizM.Name = quiz.Name
	}
	if quiz.Duration != 0 {
		quizM.Duration = quiz.Duration
	}

	if quiz.StartAt != "" {
		start_at, err := time.Parse(time.RFC3339, quiz.StartAt)
		if err != nil {
			return nil, err
		}
		quizM.StartAt = start_at
	}

	if quiz.EndAt != "" {
		endAt, err := time.Parse(time.RFC3339, quiz.EndAt)
		if err != nil {
			return nil, err
		}
		quizM.EndAt = endAt
	}

	if quiz.EndAt != "" || quiz.StartAt != "" {
		valid, err := s.quizRepo.CheckQuizDate(quizM.StartAt.Format(time.RFC3339), quizM.EndAt.Format(time.RFC3339))
		if err != nil {
			return nil, err
		}

		if !valid {
			return nil, dto.ErrQuizDateNotValid
		}
	}

	if err = s.quizRepo.Update(quizM); err != nil {
		return nil, err
	}

	return quizM, nil
}

func (s *quizService) DeleteQuiz(id uint) error {
	quizM, err := s.quizRepo.GetByID(id)
	if err != nil {
		return err
	}
	return s.quizRepo.Delete(quizM.ID)
}
