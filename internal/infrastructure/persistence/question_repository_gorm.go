package persistence

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
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

func (r *questionRepo) GetAllByFilter(filter dto.QuestionFilter) ([]*models.Question, int64, error) {
	var questions []*models.Question
	var total int64

	offset, limit := (filter.Page-1)*filter.PageSize, filter.PageSize
	query := r.db.Model(&models.Question{}).Where("quiz_id = ?", filter.QuizID).Preload("Answers").Count(&total)
	if err := query.Offset(offset).Limit(limit).Find(&questions).Error; err != nil {
		return nil, 0, err
	}

	return questions, total, nil
}

func (r *questionRepo) TotalQuestions(quizID uint) (int64, error) {
	var total int64
	err := r.db.Model(&models.Question{}).Where("quiz_id = ?", quizID).Preload("Answers").Count(&total).Error
	return total, err
}
