package dto

type UserFilter struct {
	Email string `query:"email"`
}

type UserLogin struct {
	Email    string `body:"email" validate:"required,email"`
	Password string `body:"password" validate:"required"`
}

type UserRegister struct {
	Email     string `body:"email" validate:"required,email"`
	Password  string `body:"password" validate:"required"`
	FirstName string `body:"first_name" validate:"required"`
	LastName  string `body:"last_name" validate:"required"`
}
