package dto

type ResultCreate struct {
	QuizID uint `validate:"required"`
	UserID uint `validate:""`
}

type ResultFilter struct {
	QuizID uint
	UserID uint
	PageFilter
}
