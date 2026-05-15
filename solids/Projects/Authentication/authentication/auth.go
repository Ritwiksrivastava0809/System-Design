package auth

import (
	"fmt"
	interfaces "system-design/solids/Authentication/interface"
)

// ================================================================
// Authentication Service
// ================================================================
//
// High-level module.
//
// Depends on abstraction,
// NOT concrete repository.
//

type AuthenticationService struct {
	userRepo interfaces.UserGetter
}

// Constructor Injection
func NewAuthenticationService(
	userRepo interfaces.UserGetter,
) *AuthenticationService {

	return &AuthenticationService{
		userRepo: userRepo,
	}
}

// AuthenticateUser validates credentials.
func (service *AuthenticationService) AuthenticateUser(
	username,
	password string,
) (bool, error) {

	user, err :=
		service.userRepo.GetUser(username)

	if err != nil {
		return false, err
	}

	if user.Password != password {
		return false,
			fmt.Errorf(
				"invalid credentials",
			)
	}

	return true, nil
}
