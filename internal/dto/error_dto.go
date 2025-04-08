package dto

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

var ErrEmailExists = errors.New("email already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUserNotFound = errors.New("user not found")

var ErrMultipleCorrectAnswers = errors.New("you must set one correct answer")
var ErrNoAnswersProvided = errors.New("you must provide at least 4 answer")
