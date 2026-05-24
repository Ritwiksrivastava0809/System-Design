package authentication

import (
	"gatekeeper/constants"
	"gatekeeper/internal/shared/utils"
	"gatekeeper/internal/user"
	"gatekeeper/platform/hashing"

	"github.com/rs/zerolog"
)

type Service struct {
	userRepo  user.UserRepository
	tokenRepo TokenManager
	hasher    hashing.Hasher
	log       zerolog.Logger
}

func NewAuthService(userRepo user.UserRepository, Tokenepo TokenManager, hasher hashing.Hasher, log zerolog.Logger) *Service {
	return &Service{
		userRepo:  userRepo,
		tokenRepo: Tokenepo,
		hasher:    hasher,
		log:       log,
	}
}

func (s *Service) Login(req LoginRequest) (*LoginResponse, error) {
	//get user by username
	user, err := s.userRepo.GetUserByUserName(req.Username)
	if err != nil {
		s.log.Warn().Str(constants.UserName, req.Username).Err(err).Msg(constants.LogUserDoesNotExist)
		return nil, ErrUserNameDoesNotExist
	}
	//validate password
	err = s.hasher.Compare(user.PasswordHash, req.Password)
	if err != nil {
		s.log.Warn().Str(constants.UserName, req.Username).Err(err).Msg(constants.LogInvalidPassword)
		return nil, ErrInvalidPassword
	}
	//generate Token

	accessToken, err := s.tokenRepo.CreateToken(user.Username, utils.GetAccessTokenDuration(), user.Role)

	if err != nil {
		s.log.Warn().Str(constants.UserName, req.Username).Err(err).Msg(constants.LogTokenkErr)
		return nil, ErrTokenGeneration
	}

	UserResponse := UserResponse{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		UserName:    user.Username,
		Address:     *user.Address,
		DateOfBirth: user.DateOfBirth,
		Role:        string(user.Role),
	}

	//return response
	res := &LoginResponse{
		AccessToken: accessToken,
		User:        UserResponse,
	}
	return res, nil
}

func (s *Service) Logout(accessToken string) error {
	if accessToken == "" {
		s.log.Warn().Msg("logout failed: missing access token")
		return ErrAuthorizationTokenMissing
	}

	err := s.tokenRepo.RevokeToken(accessToken)
	if err != nil {
		s.log.Warn().Err(err).Msg("logout failed")
		return err
	}

	return nil
}
