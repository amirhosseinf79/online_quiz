package service

import (
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
	if err := s.quizRepo.Create(quizModel); err != nil {
		return nil, err
	}
	return quizModel, nil
}
