package authentication

import (
	"gatekeeper/internal/user"
	"time"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

type UserResponse struct {
	ID          uuid.UUID    `json:"id"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Email       string       `json:"email"`
	UserName    string       `json:"username"`
	DateOfBirth *time.Time   `json:"date_of_birth"`
	Address     user.Address `json:"address"`
	Role        string       `json:"role"`
}

type Claims struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	Role      user.Role `json:"role"`
}
