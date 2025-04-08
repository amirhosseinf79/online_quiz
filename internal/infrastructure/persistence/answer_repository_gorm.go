package persistence

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"gorm.io/gorm"
)

type answerRepo struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) repository.AnswerRepository {
	return &answerRepo{db: db}
}

func (r *answerRepo) Update(answer *models.Answer) error {
	return r.db.Save(answer).Error
}
