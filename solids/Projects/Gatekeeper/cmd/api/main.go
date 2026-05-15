package main

import (
	"flag"
	"fmt"
	"gatekeeper/config"
	"gatekeeper/platform/database"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	environment := flag.String("e", "development", "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()

	if *environment != "development" && *environment != "production" {
		fmt.Println("Invalid environment. Allowed values are 'development' and 'production'")
		os.Exit(1)
	}

	fmt.Printf("Starting server in %s environment...\n", *environment)

	// Initialize configuration
	config.Init(*environment)

	log.Info().Msg(fmt.Sprintf("Configuration loaded for %s environment", *environment))

	// Initialize database connection
	dbConnection, err := database.NewPostgresConnection()
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database" + err.Error())
		panic("can't initialize db " + err.Error())
	}
	log.Info().Msg("Database connection initialized successfully")

	err = dbConnection.Migrate(*environment == "development")
	if err != nil {
		log.Error().Err(err).Msg("Failed to migrate database" + err.Error())
		panic("can't migrate db " + err.Error())
	}
	log.Info().Msg("Database migration completed successfully")
}
