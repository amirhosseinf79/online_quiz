package dto

type QuizFilter struct {
	Name string `json:"name" query:"name"`
	PageFilter
}

type QuizCreate struct {
	Name string `json:"name" validate:"required"`
}
