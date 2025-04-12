package models

type UserAnswer struct {
	ID           uint     `gorm:"primaryKey"`
	UserResultID uint     ``
	QuestionID   uint     ``
	AnswerID     uint     ``
	Question     Question `gorm:"constraint:OnDelete:SET NULL"`
	Answer       Answer   `gorm:"constraint:OnDelete:SET NULL"`
}
