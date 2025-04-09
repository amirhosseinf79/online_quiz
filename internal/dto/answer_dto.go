package dto

type AnswerFields struct {
	Text      string `json:"text" validate:"required"`
	IsCorrect bool   `json:"is_correct"`
}

type AnswerUpdate struct {
	ID        uint   `json:"id" validate:"required" form:"id"`
	Text      string `json:"text" validate:"required" form:"text"`
	IsCorrect bool   `json:"is_correct" validate:"" form:"is_correct"`
}
