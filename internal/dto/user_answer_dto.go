package dto

type AddUserAnswer struct {
	ResultID   uint `validate:"required" body:"result_id" form:"result_id"`
	QuestionID uint `validate:"required" body:"question_id" form:"question_id"`
	AnswerID   uint `validate:"required" body:"answer_id" form:"answer_id"`
}
