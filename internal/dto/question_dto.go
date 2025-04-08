package dto

type QuestionFilter struct {
	QuizID uint `json:"quiz_id" query:"quiz_id"`
	PageFilter
}

type QuestionCreate struct {
	QuizID  uint           `json:"quiz_id" validate:"required"`
	Text    string         `json:"text" validate:"required"`
	Answers []AnswerFields `json:"answers" validate:"required,dive"`
}
