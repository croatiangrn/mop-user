package mop_user

import "errors"

var (
	ErrFirstNameBlank   = errors.New("first_name_cannot_be_blank")
	ErrLastNameBlank    = errors.New("last_name_cannot_be_blank")
	ErrEmailInvalid     = errors.New("email_is_not_valid")
	ErrPasswordTooShort = errors.New("password_too_short")
	ErrPasswordMismatch = errors.New("password_and_password_repeat_mismatch")
)
