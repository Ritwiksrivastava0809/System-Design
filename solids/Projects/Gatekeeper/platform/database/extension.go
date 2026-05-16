package database

import "gorm.io/gorm"

func EnableExtensions(db *gorm.DB) error {
	return db.Exec(`
		CREATE EXTENSION IF NOT EXISTS "pgcrypto";
	`).Error
}
