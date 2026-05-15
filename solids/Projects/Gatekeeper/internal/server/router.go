package server

import (
	"gatekeeper/constants"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	// Define server fields here, such as router, database connection, etc.
	db     *gorm.DB
	router *gin.Engine
}

func NewServer(dbConnection *gorm.DB) (*Server, error) {
	server := &Server{
		db: dbConnection,
	}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set(constants.ConstantDB, dbConnection)
		c.Next()

	})

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
	}

	server.router = router

	return server, nil
}
