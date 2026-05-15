package user

import (
	"fmt"

	"system-design/solids/Authentication/models"
)

// ================================================================
// InMemoryUserRepository
// ================================================================
//
// Concrete repository implementation.
//
// Could later be replaced with:
// - PostgresUserRepository
// - MongoUserRepository
// - RedisUserRepository
//

type InMemoryUserRepository struct {
	users []models.User
}

type PostgresUserRepository struct {
	users []models.User
}

// Constructor
func NewUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make([]models.User, 0),
	}
}

func NewPostgresUserRepository() *PostgresUserRepository {
	return &PostgresUserRepository{
		users: make([]models.User, 0),
	}
}

// AddUser stores a new user.
func (repo *InMemoryUserRepository) AddUser(
	user models.User,
) {

	repo.users = append(
		repo.users,
		user,
	)
}

func (repo *PostgresUserRepository) AddUser(
	user models.User,
) {

	repo.users = append(
		repo.users,
		user,
	)
}

// GetUser retrieves user by username.
func (repo *InMemoryUserRepository) GetUser(
	username string,
) (*models.User, error) {

	for _, user := range repo.users {

		if user.UserName == username {
			return &user, nil
		}
	}

	return nil,
		fmt.Errorf(
			"user not found: %s",
			username,
		)
}

func (repo *PostgresUserRepository) GetUser(
	username string,
) (*models.User, error) {

	for _, user := range repo.users {

		if user.UserName == username {
			return &user, nil
		}
	}

	return nil,
		fmt.Errorf(
			"user not found: %s",
			username,
		)
}
