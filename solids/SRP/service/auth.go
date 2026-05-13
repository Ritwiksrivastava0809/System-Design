package service

import (
	"fmt"
	"system-design/solids/SRP/models"
)

// AuthenticationService is responsible for handling user authentication operations.
type AuthenticationService struct {
	UserRepositry *UserRepositry
}

// Auth service is responsible for handling user authentication and authorization.
func NewAuthenticationService(userRepo *UserRepositry) *AuthenticationService {
	return &AuthenticationService{
		UserRepositry: userRepo,
	}
}

// Authenticate verifies the user's credentials and returns true if they are valid.
func (as *AuthenticationService) Authenticate(username, password string) (bool, error) {
	// Authentication logic here
	user, err := as.GetUserbyUsername(username)
	if err != nil {
		return false, err
	}
	if user.UserName == username && user.Password == password {
		return true, nil
	}
	return false, nil

}

func (as *AuthenticationService) GetUserbyUsername(username string) (models.User, error) {
	// Logic to retrieve user by username from the repository
	// For demonstration, we will return a dummy user
	for _, user := range as.UserRepositry.Users {
		if user.UserName == username {
			return user, nil
		}
	}
	return models.User{}, fmt.Errorf("user not found")
}
