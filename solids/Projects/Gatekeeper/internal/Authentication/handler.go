package authentication

import (
	"net/http"
	"strings"

	"gatekeeper/constants"
	"gatekeeper/constants/errorlogs"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AuthHandler struct {
	service *Service
}

func NewAuthHandler(service *Service) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (s *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Ctx(c).Warn().Err(err).Msg(errorlogs.BindJsonError)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
		})
		return
	}

	resp, err := s.service.Login(req)
	if err != nil {
		s.switchLoginError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, resp)
}

func (s *AuthHandler) Logout(c *gin.Context) {
	authorizationHeader := c.GetHeader(constants.Authorization)
	if authorizationHeader == "" {
		log.Ctx(c).Warn().Msg("missing authorization header")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": ErrAuthorizationTokenMissing.Error()})
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authorizationHeader, bearerPrefix) {
		log.Ctx(c).Warn().Msg("malformed authorization header")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": ErrAuthorizationTokenMalformed.Error()})
		return
	}

	accessToken := strings.TrimPrefix(authorizationHeader, bearerPrefix)
	if err := s.service.Logout(accessToken); err != nil {
		log.Ctx(c).Warn().Err(err).Msg("logout failed")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

func (s *AuthHandler) switchLoginError(c *gin.Context, err error) {
	switch err {
	case ErrUserNameDoesNotExist, ErrInvalidPassword:
		log.Ctx(c).Warn().Err(err).Msg(constants.LogInvalidPassword)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	case ErrTokenGeneration:
		log.Ctx(c).Error().Err(err).Msg(constants.LogTokenkErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate authentication token"})
	default:
		log.Ctx(c).Error().Err(err).Msg("authentication login failed")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to authenticate user"})
	}
}
