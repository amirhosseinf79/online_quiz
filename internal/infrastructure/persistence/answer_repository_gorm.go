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

func (r *answerRepo) GetByID(id uint) (model *models.Answer, err error) {
	err = r.db.First(&model, id).Error
	return
}

func (r *answerRepo) Update(answer *models.Answer) error {
	return r.db.Save(answer).Error
}

func (r *answerRepo) CountCorrectAnswers(questionID, excludeID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Answer{}).
		Where("question_id = ? AND id != ? AND is_correct = ?", questionID, excludeID, true).
		Count(&count).Error
	return count, err
}

func (r *answerRepo) SetOtherAnswersIncorrect(questionID, excludeID uint) error {
	return r.db.Model(&models.Answer{}).
		Where("question_id = ? AND id != ? AND is_correct = ?", questionID, excludeID, true).
		Updates(map[string]any{"is_correct": false}).Error
}
