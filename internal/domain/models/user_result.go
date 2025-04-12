package models

type UserResult struct {
	ID      uint         `gorm:"primaryKey"`
	QuizID  uint         ``
	UserID  uint         ``
	Score   float64      ``
	Quiz    Quiz         `gorm:"constraint:OnDelete:SET NULL"`
	User    User         `gorm:"constraint:OnDelete:SET NULL"`
	Answers []UserAnswer `gorm:"foreignKey:UserResultID;constraint:OnDelete:SET NULL"`
}
