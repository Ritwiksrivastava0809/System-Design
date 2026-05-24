package config

import (
	"fmt"
	"gatekeeper/constants"
	"gatekeeper/constants/errorlogs"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {

	var err error
	config = viper.New()

	config.SetConfigType(constants.DefaultConfigurationType)
	config.SetConfigName(env)
	config.AddConfigPath(constants.DefaultConfigurationPath)

	err = config.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf(errorlogs.ParsingError, err.Error()))
	}
}

// DatabaseConfig returns the database configuration from the config file
func DatabaseConfig() *PostgresConfig {
	return &PostgresConfig{
		UserName: config.GetString("database.username"),
		Password: config.GetString("database.password"),
		Host:     config.GetString("database.host"),
		Port:     config.GetInt("database.port"),
		DBName:   config.GetString("database.name"),
		SSLMode:  config.GetString("database.sslmode"),
	}
}

func GetConfig() *viper.Viper {
	return config
}

func GetInternalToken() string {
	return config.GetString("token.internal")
}

func GetSymmetricKey() string {
	return config.GetString("token.symmetric")
}

func GetAccessTokenDuration() string {
	return config.GetString("token.access.duration")
}
