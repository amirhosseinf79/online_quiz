package models

type Answer struct {
	ID         uint   `gorm:"primaryKey; not null" json:"id"`
	QuestionID uint   `gorm:"not null" json:"-"`
	Text       string `gorm:"not null" json:"text"`
	IsCorrect  bool   `gorm:"not null" json:"-"`
}
