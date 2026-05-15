package server

import (
	"fmt"
	"gatekeeper/config"
	"gatekeeper/constants/errorlogs"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func Init(dbConnection *gorm.DB) {
	config := config.GetConfig()
	// Initialize other server components here, such as routes, middleware, etc.

	router, err := NewServer(dbConnection)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf(errorlogs.ServerError, err))
	}

	serverURL := fmt.Sprintf("%s:%d", config.GetString("server.host"), config.GetInt("server.port"))
	log.Info().Msg(fmt.Sprintf("Starting server at %s", serverURL))
	if err := router.router.Run(serverURL); err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf(errorlogs.ServerError, err))
	}

}
