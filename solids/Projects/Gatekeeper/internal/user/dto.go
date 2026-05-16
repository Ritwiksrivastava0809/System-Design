package user

type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`

	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`

	Password string `json:"password" binding:"required,min=8"`

	DateOfBirth *string `json:"date_of_birth,omitempty"`

	Role *Role `json:"role,omitempty"`

	Address *AddressRequest `json:"address,omitempty"`
}

type AddressRequest struct {
	Street  string `json:"street,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	ZipCode string `json:"zip_code,omitempty"`
	Country string `json:"country,omitempty"`
}

type CreateUserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Email    string `json:"email"`
	Username string `json:"username"`

	DateOfBirth *string `json:"date_of_birth,omitempty"`

	Role Role `json:"role"`

	Address *AddressResponse `json:"address,omitempty"`
}

type AddressResponse struct {
	Street  string `json:"street,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	ZipCode string `json:"zip_code,omitempty"`
	Country string `json:"country,omitempty"`
}
