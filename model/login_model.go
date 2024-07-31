package model

type RegisterRequest struct {
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	PhoneNumber          string `json:"phone_number" validate:"required,number,required"`
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=8"`
	PIN                  string `json:"pin" validate:"required,number,required,min=6,max=6"`
	Address              string `json:"address" validate:"required"`
	Province             string `json:"province" validate:"required"`
	City                 string `json:"city" validate:"required"`
	PostalCode           string `json:"postal_code" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
