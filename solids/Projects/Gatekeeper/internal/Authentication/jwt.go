package authentication

import (
	"fmt"
	"gatekeeper/constants"
	"gatekeeper/internal/user"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	secretKey     string
	revokedTokens map[string]time.Time
	revokedMu     sync.RWMutex
}

func NewTokenService(secretKey string) *TokenService {
	return &TokenService{
		secretKey:     secretKey,
		revokedTokens: make(map[string]time.Time),
	}
}

// CreateToken creates a new token for a specific username and duration
func (maker *TokenService) CreateToken(username string, duration time.Duration, role user.Role) (string, error) {
	claims, err := NewClaims(username, duration, role)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if the token is valid or not
func (maker *TokenService) parseToken(token string) (*Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf(constants.InvalidToken)
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && verr.Errors == constants.JWTValidationErrorExpired {
			return nil, fmt.Errorf(constants.ExipredToken)
		}
		return nil, fmt.Errorf(constants.InvalidToken)
	}
	payload, ok := jwtToken.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf(constants.InvalidToken)
	}

	return payload, nil
}

func (maker *TokenService) VerifyToken(token string) (*Claims, error) {
	claims, err := maker.parseToken(token)
	if err != nil {
		return nil, err
	}

	if maker.isTokenRevoked(token) {
		return nil, ErrTokenRevoked
	}

	return claims, nil
}

func (maker *TokenService) RevokeToken(token string) error {
	claims, err := maker.parseToken(token)
	if err != nil {
		return err
	}

	maker.revokedMu.Lock()
	maker.revokedTokens[token] = claims.ExpiredAt
	maker.revokedMu.Unlock()

	return nil
}

func (maker *TokenService) isTokenRevoked(token string) bool {
	maker.cleanupRevokedTokens()

	maker.revokedMu.RLock()
	defer maker.revokedMu.RUnlock()

	_, revoked := maker.revokedTokens[token]
	return revoked
}

func (maker *TokenService) cleanupRevokedTokens() {
	maker.revokedMu.Lock()
	defer maker.revokedMu.Unlock()

	now := time.Now()
	for token, expiry := range maker.revokedTokens {
		if now.After(expiry) {
			delete(maker.revokedTokens, token)
		}
	}
}
