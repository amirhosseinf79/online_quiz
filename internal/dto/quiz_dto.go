package dto

type QuizFilter struct {
	Name    string `form:"name" query:"name"`
	StartAt string `form:"start_at" validate:""`
	EndAt   string `form:"end_at" validate:""`
	PageFilter
}

type QuizDate struct {
	StartAt string `form:"start_at" validate:"required"`
	EndAt   string `form:"end_at" validate:"required"`
}

type QuizCreate struct {
	Name     string `form:"name" validate:"required"`
	Duration int    `form:"duration" validate:"required"`
	QuizDate
}

type QuizUpdate struct {
	ID       uint   `form:"id" validate:"required"`
	Name     string `form:"name" validate:"required"`
	Duration int    `form:"duration" validate:"required"`
	StartAt  string `form:"start_at" validate:""`
	EndAt    string `form:"end_at" validate:""`
}
