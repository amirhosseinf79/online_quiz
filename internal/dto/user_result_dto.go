package dto

type ResultCreate struct {
	QuizID uint `validate:"required" query:"quiz_id"`
	UserID uint `validate:""`
}

type ResultFilter struct {
	QuizID uint
	UserID uint
	PageFilter
}
