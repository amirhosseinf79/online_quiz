package persistence

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userResultRepo struct {
	db *gorm.DB
}

func NewUserResultRepo(db *gorm.DB) repository.UserResultRepository {
	return &userResultRepo{db: db}
}

func (r *userResultRepo) Create(result *models.UserResult) error {
	return r.db.Create(result).Error
}

func (r *userResultRepo) Update(result *models.UserResult) error {
	return r.db.Save(result).Error
}

func (r *userResultRepo) UpdateScore(id uint, score float64) error {
	err := r.db.Model(models.UserResult{}).Omit(clause.Associations).Where("id = ?", id).Update("score", score).Error
	return err
}

func (r *userResultRepo) GetByID(id uint) (*models.UserResult, error) {
	var result models.UserResult
	err := r.db.Preload("Answers").Preload("Quiz").First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userResultRepo) GetByQuizUSerID(quizID, userID uint) (*models.UserResult, error) {
	var result models.UserResult
	err := r.db.Preload("Answers").Preload("Quiz").Where("quiz_id = ? AND user_id = ?", quizID, userID).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userResultRepo) GetAll(filter dto.ResultFilter) ([]*models.UserResult, int64, error) {
	var results []*models.UserResult
	model := r.db.Model(&models.UserResult{})
	var total int64

	if filter.QuizID > 0 {
		model.Where("quiz_id = ?", filter.QuizID)
	}
	if filter.UserID > 0 {
		model.Where("user_id = ?", filter.UserID)
	}
	limit := filter.PageSize
	offset := (filter.Page - 1) * limit
	err := model.Count(&total).Offset(offset).Limit(limit).Find(&results).Error
	if err != nil {
		return nil, total, err
	}
	return results, total, nil
}

func (r *userResultRepo) ResultExists(quizID, userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.UserResult{}).Where("quiz_id = ? AND user_id = ?", quizID, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
