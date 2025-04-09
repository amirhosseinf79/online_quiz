package dto

type RefreshToken struct {
	RefreshToken string `form:"refresh_token" validate:"required"`
}
