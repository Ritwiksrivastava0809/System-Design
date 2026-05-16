package user

import "errors"

var (
	ErrFirstNameRequired         = errors.New("first_name is required")
	ErrLastNameRequired          = errors.New("last_name is required")
	ErrEmailRequired             = errors.New("email is required")
	ErrInvalidEmailFormat        = errors.New("invalid email format")
	ErrUsernameRequired          = errors.New("username is required")
	ErrPasswordRequired          = errors.New("password is required")
	ErrPasswordTooShort          = errors.New("password must be at least 10 characters long")
	ErrPasswordInvalidCharacters = errors.New("password must include both letters and numbers")
	ErrUserEmailAlreadyExists    = errors.New("user with this email already exists")
	ErrUserUsernameAlreadyExists = errors.New("user with this username already exists")
	ErrEmailAlreadyExists        = errors.New("email already exists")
)
