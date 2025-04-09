package dto

type QuizFilter struct {
	Name string `body:"name" query:"name"`
	QuizDate
	PageFilter
}

type QuizDate struct {
	StartAt string `body:"start_at" validate:"required"`
	EndAt   string `body:"end_at" validate:"required"`
}

type QuizCreate struct {
	Name     string `body:"name" validate:"required"`
	Duration int    `body:"duration" validate:"required"`
	QuizDate
}

type QuizUpdate struct {
	ID       uint   `body:"id" validate:"required"`
	Name     string `body:"name" validate:"required"`
	Duration int    `body:"duration" validate:"required"`
	QuizDate
}
