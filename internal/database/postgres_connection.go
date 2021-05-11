package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // we need that for swagger documentation

	"github.com/pcherednichenko/users/internal/config"
	"github.com/pcherednichenko/users/internal/models"
)

const notFoundErr = "user with id %d not found"

// PostgresDB contains basically connection to postgres with gorm to run queries
type PostgresDB struct {
	db *gorm.DB
}

// NewPostgresConnection creates connection to postgres database and checks it
func NewPostgresConnection(c config.Config) (*PostgresDB, error) {
	connectionInfo := fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=disable",
		c.User, c.Password, c.Host, c.DBName,
	)
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	// TODO: instead of this migration you should use proper migrations in separate files
	err = db.AutoMigrate(&models.User{}).Error
	if err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}
	// checking database connection
	err = db.Exec("SELECT 1").Error
	if err != nil {
		return nil, fmt.Errorf("can not connect to db: %w", err)
	}

	return &PostgresDB{
		db: db,
	}, nil
}
