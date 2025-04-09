package persistence

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"github.com/amirhosseinf79/online_quiz/internal/dto"
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
	if answer.IsCorrect {
		model := r.db.Model(&models.Answer{})
		model = model.Where("question_id = ? AND id != ? AND is_correct = ?", answer.QuestionID, answer.ID, true)
		err := model.Updates(map[string]any{"is_correct": false}).Error
		if err != nil {
			return err
		}
	} else {
		model := r.db.Model(&models.Answer{})
		model = model.Where("question_id = ? AND id != ? AND is_correct = ?", answer.QuestionID, answer.ID, true)
		var count int64
		model.Count(&count)
		if count == 0 {
			return dto.ErrMultipleCorrectAnswers
		}
	}
	return r.db.Save(answer).Error
}
