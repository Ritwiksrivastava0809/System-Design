package server

import (
	"gatekeeper/constants"
	"gatekeeper/internal/user"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
	router *gin.Engine
}

func NewServer(userHandler *user.UserHandler) (*Server, error) {
	server := &Server{
		server: &http.Server{},
	}

	router := gin.Default()

	// router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions}
	corsConfig.AllowHeaders = []string{constants.Origin, constants.ContentType, constants.ContentLength, constants.Authorization}

	router.Use(cors.New(corsConfig))
	v0 := router.Group("/api/v0")
	{
		v0.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
		// Define other routes here

		userGroup := v0.Group("/users")
		{
			userGroup.POST("/", userHandler.CreateUser)
			// Define other user-related routes here
		}

	}

	server.router = router

	return server, nil
}
