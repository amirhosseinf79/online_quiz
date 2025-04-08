package persistence

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"gorm.io/gorm"
)

type questionRepo struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) repository.QuestionRepository {
	return &questionRepo{db: db}
}

func (r *questionRepo) Create(question *models.Question) error {
	return r.db.Create(question).Error
}

func (r *questionRepo) Update(question *models.Question) error {
	return r.db.Save(question).Error
}

func (r *questionRepo) Delete(id uint) error {
	return r.db.Delete(&models.Question{}, id).Error
}

func (r *questionRepo) GetByID(id uint) (*models.Question, error) {
	var question models.Question
	if err := r.db.Preload("Answers").First(&question, id).Error; err != nil {
		return nil, err
	}
	return &question, nil
}
