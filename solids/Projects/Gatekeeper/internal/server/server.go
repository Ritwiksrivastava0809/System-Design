package server

import (
	"fmt"
	"gatekeeper/config"
	"gatekeeper/constants/errorlogs"
	"gatekeeper/internal/user"

	"github.com/rs/zerolog/log"
)

func Init(userHandler *user.UserHandler) {
	config := config.GetConfig()
	// Initialize other server components here, such as routes, middleware, etc.

	router, err := NewServer(userHandler)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf(errorlogs.ServerError, err))
	}

	serverURL := fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetInt("server.port"))
	log.Info().Msg(fmt.Sprintf("Starting server at %s", serverURL))
	if err := router.router.Run(serverURL); err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf(errorlogs.ServerError, err))
	}

}
