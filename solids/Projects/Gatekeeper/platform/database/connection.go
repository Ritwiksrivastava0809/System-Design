package database

import (
	"fmt"
	"gatekeeper/config"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnection struct {
	connection *gorm.DB
}

func NewPostgresConnection() (*PostgresConnection, error) {
	dbConfig := config.DatabaseConfig()
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.SSLMode,
	)

	log.Info().Msg("Attempting to connect to the database with DSN: " + dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true, //Disable prepared statement cache
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	//connection pool settings
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	return &PostgresConnection{
		connection: db,
	}, nil

}

func (p *PostgresConnection) GetConnection() *gorm.DB {
	return p.connection
}

func (p *PostgresConnection) HealthCheck() error {
	sqlDB, err := p.connection.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
