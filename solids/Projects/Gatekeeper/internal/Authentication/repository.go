package authentication

import (
	"gatekeeper/internal/user"
	"time"
)

type AuthRepository interface {
	Login(LoginRequest) (*user.User, error)
	Logout() error
}

type TokenManager interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration, role user.Role) (string, error)

	// VerifyToken Checks if the token is valid or not
	VerifyToken(token string) (*Claims, error)

	// RevokeToken invalidates a token until it expires
	RevokeToken(token string) error
}
