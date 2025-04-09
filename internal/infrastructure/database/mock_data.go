package database

import (
	"log"
	"time"

	"github.com/amirhosseinf79/online_quiz/internal/domain/models"
	"gorm.io/gorm"
)

func AddMockData(db *gorm.DB) {
	// Mock Users
	users := []models.User{
		{FirstName: "Admin", LastName: "User", Email: "admin@example.com", Password: "$2a$10$hashedpassword", Role: "admin"},
		{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Password: "$2a$10$hashedpassword", Role: "user"},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Printf("Error adding user: %v", err)
		}
	}

	// Mock Quizzes
	quizzes := []models.Quiz{
		{Name: "Math Quiz", StartAt: time.Now().Add(-1 * time.Hour), EndAt: time.Now().Add(1 * time.Hour), Duration: 60},
		{Name: "Science Quiz", StartAt: time.Now().Add(-2 * time.Hour), EndAt: time.Now().Add(2 * time.Hour), Duration: 120},
	}

	for _, quiz := range quizzes {
		if err := db.Create(&quiz).Error; err != nil {
			log.Printf("Error adding quiz: %v", err)
		}
	}

	// Mock Questions
	questions := []models.Question{
		{QuizID: 1, Question: "What is 2 + 2?"},
		{QuizID: 1, Question: "What is the square root of 16?"},
		{QuizID: 2, Question: "What is the chemical symbol for water?"},
	}

	for _, question := range questions {
		if err := db.Create(&question).Error; err != nil {
			log.Printf("Error adding question: %v", err)
		}
	}

	// Mock Answers
	answers := []models.Answer{
		{QuestionID: 1, Text: "4", IsCorrect: true},
		{QuestionID: 1, Text: "5", IsCorrect: false},
		{QuestionID: 2, Text: "4", IsCorrect: true},
		{QuestionID: 2, Text: "5", IsCorrect: false},
		{QuestionID: 3, Text: "H2O", IsCorrect: true},
		{QuestionID: 3, Text: "O2", IsCorrect: false},
	}

	for _, answer := range answers {
		if err := db.Create(&answer).Error; err != nil {
			log.Printf("Error adding answer: %v", err)
		}
	}

	log.Println("Mock data added successfully!")
}
