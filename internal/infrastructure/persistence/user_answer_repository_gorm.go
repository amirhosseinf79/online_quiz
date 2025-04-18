package persistence

import (
	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"github.com/amirhosseinf79/online_quiz/internal/domain/repository"
	"gorm.io/gorm"
)

type userAnswerRepo struct {
	db *gorm.DB
}

func NewUserAnswerRepo(db *gorm.DB) repository.UserAnswerRepository {
	return &userAnswerRepo{db: db}
}

func (r *userAnswerRepo) Create(answer *models.UserAnswer) error {
	return r.db.Create(answer).Error
}

func (r *userAnswerRepo) Update(answer *models.UserAnswer) error {
	return r.db.Save(answer).Error
}

func (r *userAnswerRepo) GetAnswer(resultID, questionID uint) (*models.UserAnswer, error) {
	var answer *models.UserAnswer
	err := r.db.Model(&models.UserAnswer{}).Where("user_result_id = ? AND question_id = ?", resultID, questionID).First(&answer).Error
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func (r *userAnswerRepo) GetCurrectAnsCount(resultID uint) (int64, error) {
	var trueAnswers int64
	model := r.db.Model(models.UserAnswer{}).Joins("JOIN answers ON answers.id = answer_id")
	model.Where("user_result_id = ? AND answers.is_correct = true", resultID).Count(&trueAnswers)
	err := model.Error
	return trueAnswers, err
}
