package user

import (
	"net/mail"
	"regexp"
	"unicode"
)

func (s *Service) ValidateCreateUser(req CreateUserRequest) error {
	if req.FirstName == "" {
		return ErrFirstNameRequired
	}
	if req.LastName == "" {
		return ErrLastNameRequired
	}
	if req.Email == "" {
		return ErrEmailRequired
	}
	// Validate email format: must be valid RFC 5322 and match standard email pattern
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return ErrInvalidEmailFormat
	}
	// Additional regex check to ensure proper email format (local@domain.tld)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		return ErrInvalidEmailFormat
	}
	if req.Username == "" {
		return ErrUsernameRequired
	}
	if req.Password == "" {
		return ErrPasswordRequired
	}
	if len(req.Password) < 10 {
		return ErrPasswordTooShort
	}

	var hasLetter, hasNumber bool
	for _, r := range req.Password {
		if unicode.IsLetter(r) {
			hasLetter = true
		}
		if unicode.IsNumber(r) {
			hasNumber = true
		}
	}
	if !hasLetter || !hasNumber {
		return ErrPasswordInvalidCharacters
	}

	return nil
}
