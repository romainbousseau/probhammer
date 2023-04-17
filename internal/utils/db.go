package utils

import (
	"fmt"
	"testing"

	"github.com/romainbousseau/probhammer/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenDBConnection opens a db connection based on config
func OpenDBConnection(config Config) (*gorm.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.SSLMode)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  psqlconn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Migrate handles migration on the DB
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		models.Datasheet{},
		models.WeaponProfile{},
	); err != nil {
		return err
	}

	return nil
}

// SetTestStorage sets DB and storage for the test env
// It provides a cleanup function to be called in tests with:
// t.Cleanup(cleanup)
func SetTestDB(t *testing.T) (*gorm.DB, func()) {
	config, err := LoadConfig("../..", ".env.test")
	if err != nil {
		panic(err)
	}

	db, err := OpenDBConnection(config)
	if err != nil {
		panic(err)
	}

	err = Migrate(db)
	if err != nil {
		panic(err)
	}

	cleanup := func() {
		err = db.Migrator().DropTable(&models.Datasheet{})
		if err != nil {
			panic(err)
		}
	}

	return db, cleanup
}

// DropAllTables drops all tables. Use with caution!
func DropAllTables(db *gorm.DB) error {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return err
	}

	for _, table := range tables {
		err = db.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	return nil
}
