package mop_user

import "errors"

var (
	ErrFirstNameBlank   = errors.New("first_name_cannot_be_blank")
	ErrLastNameBlank    = errors.New("last_name_cannot_be_blank")
	ErrEmailInvalid     = errors.New("email_is_not_valid")
	ErrEmailTaken       = errors.New("email_already_taken")
	ErrPasswordTooShort = errors.New("password_too_short")
	ErrPasswordMismatch = errors.New("password_and_password_repeat_mismatch")
	ErrInternal         = errors.New("internal_error")
	ErrHashingPassword  = errors.New("could_not_hash_the_password")
)
