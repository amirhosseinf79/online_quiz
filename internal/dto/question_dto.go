package dto

type QuestionFilter struct {
	QuizID uint `json:"quiz_id" query:"quiz_id"`
	PageFilter
}

type QuestionCreate struct {
	QuizID  uint           `body:"quiz_id" validate:"required"`
	Text    string         `body:"text" validate:"required"`
	Answers []AnswerFields `body:"answers" validate:"required,dive"`
}

type QuestionUpdate struct {
	ID uint `body:"id" validate:"required"`
	QuestionCreate
}
