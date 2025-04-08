package repository

import "github.com/amirhosseinf79/online_quiz/internal/domain/models"

type AnswerRepository interface {
	Update(answer *models.Answer) error
}
