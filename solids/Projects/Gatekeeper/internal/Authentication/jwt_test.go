package authentication

import (
	"errors"
	"gatekeeper/constants"
	"gatekeeper/internal/shared/utils"
	"gatekeeper/internal/user"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	// Initialize a silent no-op logger for testing

	// Pass nil for repo and hasher since they aren't exercised by token generation/verification
	maker := NewTokenService(utils.RandomString(32))
	require.NotEmpty(t, maker)

	username := utils.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	// Create a token
	token, err := maker.CreateToken(username, duration, user.RoleAdmin)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Verify the token
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {

	maker := NewTokenService(utils.RandomString(32))

	token, err := maker.CreateToken(utils.RandomOwner(), -time.Minute, user.RoleUser)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, constants.ExipredToken)
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewClaims(utils.RandomOwner(), time.Minute, user.RoleAdmin)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker := NewTokenService(utils.RandomString(32))

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errors.New(constants.InvalidToken).Error())
	require.Nil(t, payload)
}

func TestRevokedJWTToken(t *testing.T) {
	maker := NewTokenService(utils.RandomString(32))
	token, err := maker.CreateToken(utils.RandomOwner(), time.Minute, user.RoleAdmin)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotNil(t, payload)

	err = maker.RevokeToken(token)
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrTokenRevoked.Error())
	require.Nil(t, payload)
}
