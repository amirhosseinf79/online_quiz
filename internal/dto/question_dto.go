package dto

type QuestionFilter struct {
	QuizID uint `json:"quiz_id" query:"quiz_id"`
	PageFilter
}

type QuestionCreate struct {
	QuizID  uint           `form:"quiz_id" validate:"required"`
	Text    string         `form:"text" validate:"required"`
	Answers []AnswerFields `form:"answers" validate:"required,dive"`
}

type QuestionUpdate struct {
	ID   uint   `form:"id" validate:"required"`
	Text string `form:"text" validate:"required"`
}
