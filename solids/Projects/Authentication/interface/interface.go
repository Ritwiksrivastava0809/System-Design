package interfaces

import "system-design/solids/Authentication/models"

// ================================================================
// Interfaces
// ================================================================
//
// Small focused interfaces following:
// - ISP
// - DIP
//

// UserGetter defines only the behavior
// required by the authentication service.
type UserGetter interface {
	GetUser(username string) (*models.User, error)
}

// UserAdder defines behavior for adding users.
type UserAdder interface {
	AddUser(user models.User)
}
