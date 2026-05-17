package user

import (
	"gatekeeper/constants"
	"gatekeeper/constants/errorlogs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserHandler struct {
	service *Service
}

func NewUserHandler(service *Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// Define handler methods for user-related HTTP endpoints here (e.g., CreateUser, GetUser, etc.)

func (h *UserHandler) CreateUser(c *gin.Context) {
	// Bind the incoming JSON request to a struct
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Ctx(c).Error().Err(err).Msg(errorlogs.BindJsonError)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
		})
		return
	}

	// call the service method to create a user
	resp, err := h.service.CreateUser(req)
	if err != nil {
		switch err {
		case ErrFirstNameRequired, ErrLastNameRequired, ErrEmailRequired, ErrInvalidEmailFormat,
			ErrUsernameRequired, ErrPasswordRequired, ErrPasswordTooShort,
			ErrPasswordInvalidCharacters:
			log.Ctx(c).Warn().Err(err).Msg(constants.LogValidationError)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case ErrEmailAlreadyExists, ErrUserEmailAlreadyExists:
			log.Ctx(c).Warn().Err(err).Msg(constants.LogConflictError)
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		default:
			log.Ctx(c).Error().Err(err).Msgf(errorlogs.CreateUserError, err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create user",
			})
			return
		}
	}

	// return pretty, indented JSON
	c.IndentedJSON(http.StatusCreated, resp)
}
