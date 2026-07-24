package requests

type CreateUserRequest struct {
	Name            string `json:"name" validate:"required,min=3,max=100"`
	Email           string `json:"email" validate:"required,email"`
	Gender          string `json:"gender" validate:"required,oneof=male female other"`
	Birthdate       string `json:"birthdate" validate:"required,datetime=2006-01-02"`
	IsActive        bool   `json:"is_active"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
