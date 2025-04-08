package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID           uint           `gorm:"primaryKey; not null" json:"-"`
	UserID       uint           `gorm:"not null" json:"-"`
	Token        string         `gorm:"not null" json:"token"`
	RefreshToken string         `gorm:"not null" json:"refresh_token"`
	User         User           `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	// t.Token = generateToken()
	// t.RefreshToken = generateRefreshToken()
	return nil
}
