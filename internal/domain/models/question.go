package models

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	ID        uint           `gorm:"primaryKey; not null" json:"id"`
	QuizID    uint           `gorm:"not null" json:"-"`
	Question  string         `gorm:"not null" json:"question"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	Answers   []Answer       `gorm:"foreignKey:QuestionID" json:"answers"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
