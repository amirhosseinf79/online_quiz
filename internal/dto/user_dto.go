package dto

type UserFilter struct {
	Email string `json:"email" query:"email"`
}
