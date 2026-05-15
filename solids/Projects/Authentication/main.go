package main

import (
	"fmt"
	user "system-design/solids/Authentication/User"
	auth "system-design/solids/Authentication/authentication"
	"system-design/solids/Authentication/models"
)

func main() {

	// ============================================================
	// Create Repository
	// ============================================================

	userRepo := user.NewUserRepository()

	postgresUserRepo := user.NewPostgresUserRepository()

	// ============================================================
	// Add Users
	// ============================================================

	userRepo.AddUser(models.User{
		UserName: "john_doe",
		Password: "password123",
	})

	userRepo.AddUser(models.User{
		UserName: "jane_doe",
		Password: "password456",
	})

	postgresUserRepo.AddUser(models.User{
		UserName: "bob_smith",
		Password: "password789",
	})

	// ============================================================
	// Create Authentication Service
	// ============================================================

	authService :=
		auth.NewAuthenticationService(
			userRepo,
		)

	postgresAuthService :=
		auth.NewAuthenticationService(
			postgresUserRepo,
		)
	// ============================================================
	// Authenticate Users
	// ============================================================

	isAuthenticated, err :=
		authService.AuthenticateUser(
			"john_doe",
			"password123",
		)

	if err != nil {

		fmt.Printf(
			"Authentication failed: %v\n",
			err,
		)

	} else if isAuthenticated {

		fmt.Println(
			"John authenticated successfully!",
		)
	}

	// ============================================================

	isAuthenticated, err =
		authService.AuthenticateUser(
			"jane_doe",
			"wrongpassword",
		)

	if err != nil {

		fmt.Printf(
			"Authentication failed: %v\n",
			err,
		)

	} else if isAuthenticated {

		fmt.Println(
			"Jane authenticated successfully!",
		)

	} else {

		fmt.Println(
			"Invalid credentials.",
		)
	}
	//===========================================================
	isAuthenticated, err =
		postgresAuthService.AuthenticateUser(
			"bob_smith",
			"password789",
		)

	if err != nil {

		fmt.Printf(
			"Authentication failed: %v\n",
			err,
		)

	} else if isAuthenticated {

		fmt.Println(
			"Bob authenticated successfully!",
		)
	}
}

// ===================================================================
// SOLID Principles Notes - Authentication System
// ===================================================================
//
// This authentication system demonstrates the SOLID principles
// through clean separation of responsibilities, abstraction-driven
// design, and dependency injection.
//
// The architecture is structured into independent layers:
//
//	main
//	  ↓
//	AuthenticationService
//	  ↓
//	UserRepository Interface
//	  ↓
//	Repository Implementations
//
// This separation creates a flexible, maintainable,
// and scalable backend architecture.
//
// ===================================================================
// 1. Single Responsibility Principle (SRP)
// ===================================================================
//
// SRP states:
//
//	A module should have only one reason to change.
//
// Each component in the system has a focused responsibility.
//
// -------------------------------------------------------------------
// User Model
// -------------------------------------------------------------------
//
//	type User struct
//
// Responsibility:
//	- represents user data only
//
// It does NOT:
//	- authenticate users
//	- manage persistence
//	- contain business logic
//
// -------------------------------------------------------------------
// AuthenticationService
// -------------------------------------------------------------------
//
//	type AuthenticationService struct
//
// Responsibility:
//	- authenticate users
//
// It only:
//	- validates credentials
//	- manages authentication flow
//
// It does NOT:
//	- store users
//	- know database details
//	- handle persistence
//
// -------------------------------------------------------------------
// Repository Implementations
// -------------------------------------------------------------------
//
//	type InMemoryUserRepository struct
//	type PostgresUserRepository struct
//
// Responsibility:
//	- user storage and retrieval
//
// Repositories handle:
//	- persistence logic
//	- data access
//
// They do NOT:
//	- authenticate users
//	- contain business rules
//
// This separation improves readability,
// maintainability, and modularity.
//
// ===================================================================
// 2. Open/Closed Principle (OCP)
// ===================================================================
//
// OCP states:
//
//	Software entities should be open for extension,
//	but closed for modification.
//
// The system allows new repository implementations
// to be added without modifying authentication logic.
//
// Example:
//
//	auth.NewAuthenticationService(userRepo)
//
// and:
//
//	auth.NewAuthenticationService(postgresUserRepo)
//
// Both work without changing AuthenticationService.
//
// Future implementations can easily be added:
//
//	- MongoUserRepository
//	- RedisUserRepository
//	- MySQLUserRepository
//
// without modifying existing business logic.
//
// This makes the system extensible and adaptable
// to infrastructure changes.
//
// ===================================================================
// 3. Liskov Substitution Principle (LSP)
// ===================================================================
//
// LSP states:
//
//	Subtypes should be replaceable
//	for their abstractions.
//
// Repository implementations:
//
//	- InMemoryUserRepository
//	- PostgresUserRepository
//
// can replace each other because they satisfy
// the same repository contract.
//
// AuthenticationService does not care:
//
//	- where data comes from
//	- how storage works internally
//
// It only depends on expected behavior.
//
// This demonstrates proper substitutability
// and abstraction-driven design.
//
// ===================================================================
// 4. Interface Segregation Principle (ISP)
// ===================================================================
//
// ISP states:
//
//	Clients should not depend on methods
//	they do not use.
//
// The interfaces are small and focused.
//
// Example:
//
//	type UserGetter interface {
//		GetUser(username string) (*models.User, error)
//	}
//
// AuthenticationService only depends on:
//
//	- GetUser()
//
// It does NOT depend on:
//
//	- AddUser()
//	- DeleteUser()
//	- UpdateUser()
//
// This creates:
//
//	- focused abstractions
//	- lower coupling
//	- easier testing
//	- clearer contracts
//
// Interfaces are designed around consumer needs.
//
// ===================================================================
// 5. Dependency Inversion Principle (DIP)
// ===================================================================
//
// DIP states:
//
//	High-level modules should not depend
//	on low-level modules.
//
//	Both should depend on abstractions.
//
// -------------------------------------------------------------------
// High-Level Module
// -------------------------------------------------------------------
//
//	type AuthenticationService struct
//
// depends on abstraction:
//
//	type UserGetter interface
//
// NOT on concrete implementations:
//
//	- InMemoryUserRepository
//	- PostgresUserRepository
//
// Dependency Injection is achieved
// through constructor injection:
//
//	auth.NewAuthenticationService(userRepo)
//
// This allows:
//
//	- swappable repository implementations
//	- decoupled business logic
//	- easier testing
//	- flexible infrastructure
//
// Business logic becomes independent
// from storage implementation details.
//
// ===================================================================
// Architectural Benefits
// ===================================================================
//
// Because of SOLID principles, the system gains:
//
// 1. Flexibility
//	Repository implementations can change independently.
//
// 2. Testability
//	Mock repositories can easily be injected.
//
// 3. Maintainability
//	Changes remain isolated to specific layers.
//
// 4. Scalability
//	New storage implementations can be added safely.
//
// 5. Loose Coupling
//	Business logic is independent from infrastructure.
//
// ===================================================================
// Most Important Insight
// ===================================================================
//
// The biggest architectural achievement is:
//
//	Authentication logic no longer depends
//	on infrastructure details.
//
// This means the system can evolve:
//
//	- memory storage → postgres
//	- postgres → mongodb
//	- local database → cloud database
//
// without rewriting core authentication logic.
//
// This is the foundation of:
//
//	- clean architecture
//	- scalable backend systems
//	- enterprise-grade software design
//
// ===================================================================
