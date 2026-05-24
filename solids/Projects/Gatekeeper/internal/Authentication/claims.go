package authentication

import (
	"fmt"
	"gatekeeper/constants"
	"gatekeeper/internal/user"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

func NewClaims(username string, duration time.Duration , role user.Role) (*Claims, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	p := &Claims{
		ID:        token,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
		Role : role,
	}

	return p, nil
}

func (payload *Claims) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return fmt.Errorf(constants.ExipredToken)
	}

	return nil
}

// PasetoMaker is a Paseto token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}
