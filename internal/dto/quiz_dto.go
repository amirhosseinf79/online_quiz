package dto

type QuizFilter struct {
	Name string `json:"name" query:"name"`
}
