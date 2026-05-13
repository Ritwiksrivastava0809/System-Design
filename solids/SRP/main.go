package main

import (
	"system-design/solids/SRP/models"
	"system-design/solids/SRP/service"
)

func main() {
	// Initialize the user repository
	userRepo := service.NewUserRepository()

	// Create the authentication service with the user repository
	authService := service.NewAuthenticationService(&userRepo)

	// Add a new user to the repository
	user, err := userRepo.AddUser(models.User{
		UserName: "john_doe",
		Password: "password1234",
	})
	if err != nil {
		panic(err)
	}

	// Authenticate the user
	isAuthenticated, err := authService.Authenticate(user.UserName, user.Password)
	if err != nil {
		panic(err)
	}

	if isAuthenticated {
		println("User authenticated successfully!")
	} else {
		println("Authentication failed.")
	}

}
