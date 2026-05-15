package database

func (p *PostgresConnection) Migrate(developmentMode bool) error {
	p.connection.AutoMigrate()
	return nil
}
