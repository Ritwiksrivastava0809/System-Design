package user

import "github.com/google/uuid"

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id uuid.UUID) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByUserName(username string) (*User, error)
}
