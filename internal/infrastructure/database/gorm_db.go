package database

import (
	"log"

	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(connStr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Quiz{},
		&models.Question{},
		&models.Answer{},
	)
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	return db
}
