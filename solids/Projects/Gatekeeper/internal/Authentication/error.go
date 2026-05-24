package authentication

import "errors"

var (
	ErrUserNameDoesNotExist        = errors.New("user with provided username does not exist")
	ErrInvalidPassword             = errors.New("invalid user credentials - provided password is incorrect")
	ErrTokenGeneration             = errors.New("got error while generating the token.")
	ErrAuthorizationTokenMissing   = errors.New("authorization token is required")
	ErrAuthorizationTokenMalformed = errors.New("authorization token is malformed")
	ErrTokenRevoked                = errors.New("token has been revoked")
)
