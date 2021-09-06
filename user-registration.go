package mop_user

type UserRegistration struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
}

func NewUserRegistration() *UserRegistration {
	return &UserRegistration{}
}

func (r *UserRegistration) Validate() error {
	if len(r.FirstName) == 0 {
		return ErrFirstNameBlank
	}

	if len(r.LastName) == 0 {
		return ErrLastNameBlank
	}

	if !isEmailValid(r.Email) {
		return ErrEmailInvalid
	}

	if len(r.Password) < PassMinLength {
		return ErrPasswordTooShort
	}

	if r.Password != r.PasswordRepeat {
		return ErrPasswordMismatch
	}

	return nil
}
