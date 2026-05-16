package hashing

import "errors"

var (
	ErrInvalidHashFormat = errors.New("invalid hash format")
	ErrPasswordMismatch  = errors.New("password mismatch")
)
