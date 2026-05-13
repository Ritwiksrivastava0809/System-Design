package service

import "system-design/solids/SRP/models"

// UserRepository defines the behavior required for persisting users.
type UserRepositry struct {
	Users []models.User
}

// Initialize the user repository with an empty slice of users.
func NewUserRepository() UserRepositry {
	return UserRepositry{
		Users: make([]models.User, 0),
	}
}

// AddUser adds a new user to the repository.
func (ur *UserRepositry) AddUser(user models.User) (models.User, error) {
	ur.Users = append(ur.Users, user)
	return user, nil
}
