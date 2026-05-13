package solids

import "context"

// SOLID Principles
//
// S — Single Responsibility Principle (SRP)
// A type, function, or module should have only one reason to change.
//
// O — Open/Closed Principle (OCP)
// Software entities should be open for extension
// but closed for modification.
//
// L — Liskov Substitution Principle (LSP)
// Implementations should be replaceable through their abstractions
// without breaking expected behavior.
//
// I — Interface Segregation Principle (ISP)
// Clients should not be forced to depend on methods they do not use.
//
// D — Dependency Inversion Principle (DIP)
// High-level modules should depend on abstractions,
// not concrete implementations.

// ----------------------------------------------------------------------
// Domain Model
// ----------------------------------------------------------------------

type User struct {
	ID    int64
	Name  string
	Email string
}

// ----------------------------------------------------------------------
// Repository Abstraction
// ----------------------------------------------------------------------

// UserRepository defines the behavior required
// for persisting users.
//
// The service layer depends on this abstraction
// instead of a concrete database implementation.
//
// This follows the Dependency Inversion Principle.
type UserRepository interface {
	Create(ctx context.Context, user User) error
}

// ----------------------------------------------------------------------
// Service Layer
// ----------------------------------------------------------------------

// UserService is a high-level business module.
//
// It depends on the UserRepository abstraction
// rather than a concrete database implementation.
//
// This makes the service:
// - easier to test
// - loosely coupled
// - flexible to future changes
type UserService struct {
	repo UserRepository
}

// NewUserService injects the repository dependency.
//
// This is constructor injection,
// which is idiomatic in Go.
func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// CreateUser handles the business workflow
// for creating a user.
//
// The service delegates persistence responsibility
// to the repository layer.
func (s *UserService) CreateUser(
	ctx context.Context,
	user User,
) error {

	// business logic can be added here
	// example:
	// - validation
	// - password hashing
	// - audit logging
	// - event publishing

	return s.repo.Create(ctx, user)
}

// ----------------------------------------------------------------------
// Concrete Repository Implementation
// ----------------------------------------------------------------------

// PostgresUserRepository is a concrete implementation
// of UserRepository.
//
// It contains database-specific logic.
type PostgresUserRepository struct {
	// db *sql.DB
	// other database connection fields
}

// Create persists the user into PostgreSQL.
//
// This implementation detail is hidden from the service layer.
func (r *PostgresUserRepository) Create(
	ctx context.Context,
	user User,
) error {

	// database insert logic here

	return nil
}

type MongoUserRepository struct {
	// db *mongo.Client
	// other database connection fields
}

func (r *MongoUserRepository) Create(
	ctx context.Context,
	user User,
) error {
	// database insert logic here

	return nil
}

// ----------------------------------------------------------------------
// Why This Design Is Better
// ----------------------------------------------------------------------

/*
BAD DESIGN (tight coupling)

func Create(ctx context.Context, user *User) error {
	return db.Create(ctx, user)
}

Problem:
- business logic directly depends on database implementation
- difficult to test
- tightly coupled
- hard to replace database/storage layer

--------------------------------------------------------

GOOD DESIGN (dependency inversion)

UserService ---> UserRepository(interface) ---> PostgreSQL

Benefits:
- loose coupling
- easier testing
- flexible architecture
- clean separation of concerns
- easier maintenance

The service depends only on behavior:

    Create(ctx, user)

It does not care:
- which database is used
- how persistence works
- whether data is stored in PostgreSQL, MongoDB, or memory

This is the core idea behind
the Dependency Inversion Principle.
*/
