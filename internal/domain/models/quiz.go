package models

import (
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	ID        uint           `gorm:"primaryKey; not null" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Duration  int            `gorm:"not null" json:"duration"`
	StartAt   time.Time      `gorm:"not null" json:"start_at"`
	EndAt     time.Time      `gorm:"not null" json:"end_at"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	Questions []Question     `gorm:"foreignKey:QuizID" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
