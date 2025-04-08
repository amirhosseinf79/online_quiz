package dto

type AnswerFields struct {
	Text      string `json:"text" validate:"required"`
	IsCorrect bool   `json:"is_correct"`
}

type AnswerUpdate struct {
	ID        uint   `json:"id" validate:"required" body:"id"`
	Text      string `json:"text" validate:"required" body:"text"`
	IsCorrect bool   `json:"is_correct" validate:"required" body:"is_correct"`
}
