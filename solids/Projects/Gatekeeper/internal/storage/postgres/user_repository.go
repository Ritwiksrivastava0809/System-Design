package postgres

import (
	"gatekeeper/internal/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(u *user.User) error {
	if err := r.db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*user.User, error) {
	var u user.User
	if err := r.db.Where("id = ?", id).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*user.User, error) {
	var u user.User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetUserByUserName(username string) (*user.User, error) {
	var u user.User
	if err := r.db.Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
