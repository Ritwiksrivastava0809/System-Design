package database

import (
	"fmt"
	"gatekeeper/constants/errorlogs"
	"gatekeeper/internal/user"
)

func (p *PostgresConnection) Migrate(developmentMode bool) error {
	if err := EnableExtensions(p.connection); err != nil {
		return fmt.Errorf(errorlogs.ExtensionError, err)
	}
	if p.connection.AutoMigrate(&user.User{}) != nil {
		return fmt.Errorf(errorlogs.MigrationError, p.connection.AutoMigrate(&user.User{}))
	}
	return nil
}
