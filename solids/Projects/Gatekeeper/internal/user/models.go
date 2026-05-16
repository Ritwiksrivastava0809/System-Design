package user

import (
	"gatekeeper/constants"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`

	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Username string `gorm:"type:varchar(100);uniqueIndex;not null"`

	PasswordHash string `gorm:"type:varchar(255);not null"`

	DateOfBirth *time.Time `gorm:"type:date"`

	Role Role `gorm:"type:varchar(20);default:'user';not null"`

	Address *Address `gorm:"embedded;embeddedPrefix:address_"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Address struct {
	Street  string `gorm:"type:varchar(255)"`
	City    string `gorm:"type:varchar(100)"`
	State   string `gorm:"type:varchar(100)"`
	ZipCode string `gorm:"type:varchar(20)"`
	Country string `gorm:"type:varchar(100)"`
}

func (User) TableName() string {
	return constants.UserTableName
}
