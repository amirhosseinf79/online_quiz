package persistence

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type quizRepo struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) repository.QuizRepository {
	return &quizRepo{db: db}
}

func (r *quizRepo) Create(quiz *models.Quiz) error {
	return r.db.Omit(clause.Associations).Create(quiz).Error
}

func (r *quizRepo) Update(quiz *models.Quiz) error {
	return r.db.Omit(clause.Associations).Save(quiz).Error
}

func (r *quizRepo) Delete(id uint) error {
	return r.db.Delete(&models.Quiz{}, id).Error
}

func (r *quizRepo) GetByID(id uint) (quiz *models.Quiz, err error) {
	if err := r.db.First(&quiz, id).Error; err != nil {
		return nil, err
	}
	return
}

func (r *quizRepo) GetAllByFilter(filter dto.QuizFilter) (quizzes []*models.Quiz, total int64, err error) {
	query := r.db.Model(&models.Quiz{})
	if filter.Name != "" {
		query = query.Where("LOWER(name) LIKE ?", "LOWER(%"+filter.Name+"%)")
	}
	offset, limit := (filter.Page-1)*filter.PageSize, filter.PageSize
	err = query.Count(&total).Offset(offset).Limit(limit).Find(&quizzes).Error
	return
}
