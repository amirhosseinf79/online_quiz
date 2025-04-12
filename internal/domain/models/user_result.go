package models

type UserResult struct {
	ID      uint         `gorm:"primaryKey"`
	QuizID  uint         `json:"-"`
	UserID  uint         `json:"-"`
	Score   float64      `json:"score"`
	Quiz    Quiz         `gorm:"constraint:OnDelete:SET NULL" json:"quiz"`
	User    User         `gorm:"constraint:OnDelete:SET NULL" json:"-"`
	Answers []UserAnswer `gorm:"foreignKey:UserResultID;constraint:OnDelete:SET NULL" json:"answers"`
}
