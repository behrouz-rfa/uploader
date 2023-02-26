package postgres

import (
	"fmt"
	"gorm.io/gorm"
	"uploader/files/internal/domain"
)

// Migrate creates the schema in the database.
func Migrate(db *gorm.DB) error {

	if err := db.AutoMigrate(&domain.FileInfo{}); err != nil {
		return fmt.Errorf("migrate ORM tables: %w", err)
	}
	return nil
}
