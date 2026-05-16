package hashing

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"gatekeeper/constants/errorlogs"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon2Hasher struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

func NewArgon2Hasher() *Argon2Hasher {
	return &Argon2Hasher{
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}
}

func (a *Argon2Hasher) Hash(password string) (string, error) {

	salt := make([]byte, 16)

	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf(errorlogs.SaltGenerationError, err)
	}

	hash := a.generateHash(password, salt)

	saltEncoded := base64.RawStdEncoding.EncodeToString(salt)

	return fmt.Sprintf("%s:%s", saltEncoded, hash), nil
}

func (a *Argon2Hasher) Compare(
	storedPassword string,
	providedPassword string,
) error {

	parts := strings.Split(storedPassword, ":")

	if len(parts) != 2 {
		return ErrInvalidHashFormat
	}

	saltEncoded := parts[0]
	expectedHash := parts[1]

	salt, err := base64.RawStdEncoding.DecodeString(saltEncoded)
	if err != nil {
		return fmt.Errorf("failed to decode salt: %w", err)
	}

	providedHash := a.generateHash(providedPassword, salt)

	if subtle.ConstantTimeCompare(
		[]byte(providedHash),
		[]byte(expectedHash),
	) != 1 {
		return ErrPasswordMismatch
	}

	return nil
}

func (a *Argon2Hasher) generateHash(
	password string,
	salt []byte,
) string {

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		a.time,
		a.memory,
		a.threads,
		a.keyLen,
	)

	return base64.RawStdEncoding.EncodeToString(hash)
}
