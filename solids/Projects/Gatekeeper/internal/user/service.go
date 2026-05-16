package user

import (
	"errors"
	"gatekeeper/platform/hashing"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	repo UserRepository
	hash hashing.Hasher
}

func NewService(userRepo UserRepository, hash hashing.Hasher) *Service {
	return &Service{
		repo: userRepo,
		hash: hash,
	}
}

func (s *Service) CreateUser(req CreateUserRequest) (*CreateUserResponse, error) {
	//validate user data here , check if email is valid, password meets criteria, etc.
	err := s.ValidateCreateUser(req)
	if err != nil {
		return nil, err
	}
	//check if user with the same email already exists
	user, err := s.repo.GetUserByEmail(req.Email)

	if err == nil && user != nil {
		return nil, ErrEmailAlreadyExists
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	//hash the password before saving to the database
	hashedPassword, err := s.hash.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	//Parse date of birth if provided
	var dob *time.Time
	if req.DateOfBirth != nil && *req.DateOfBirth != "" {
		parsed, err := time.Parse("2006-01-02", *req.DateOfBirth)
		if err != nil {
			return nil, err
		}
		dob = &parsed
	}

	var address *Address
	if req.Address != nil {
		address = &Address{
			Street:  req.Address.Street,
			City:    req.Address.City,
			State:   req.Address.State,
			ZipCode: req.Address.ZipCode,
			Country: req.Address.Country,
		}
	}

	role := RoleUser
	if req.Role != nil {
		role = *req.Role
	}

	user = &User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Address:      address,
		Role:         role,
		DateOfBirth:  dob,
		Username:     req.Username,
	}
	//save the user to the database
	err = s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	response := &CreateUserResponse{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Role:      user.Role,
	}

	if user.DateOfBirth != nil {
		dobStr := user.DateOfBirth.Format("2006-01-02")
		response.DateOfBirth = &dobStr
	}

	if user.Address != nil {
		response.Address = &AddressResponse{
			Street:  user.Address.Street,
			City:    user.Address.City,
			State:   user.Address.State,
			ZipCode: user.Address.ZipCode,
			Country: user.Address.Country,
		}
	}

	return response, nil
}

// validation moved to validation.go
