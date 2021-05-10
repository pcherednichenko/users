package config

import (
	"github.com/pcherednichenko/users/pkg/env"
)

// Config contain all service configuration
type Config struct {
	Application
	Database
}

// Application configuration
type Application struct {
	Port string
}

// Database configuration of connection
type Database struct {
	User     string
	Password string
	Host     string
	DBName   string
}

// LoadConfigFromEnv variables with default options
// TODO: load config from yaml for example
func LoadConfigFromEnv(l env.Logger) Config {
	// TODO: in general for example for database better to fail instead of using default configuration
	return Config{
		Application{
			Port: env.GetOrDefault(l, "PORT", ":8080"),
		},
		Database{
			User: env.GetOrDefault(l, "DB_USER", "test_user"),
			Password: env.GetOrDefault(l, "DB_PASS", "test_pass"),
			Host: env.GetOrDefault(l, "DB_HOST", "postgres"),
			DBName: env.GetOrDefault(l, "DB_NAME", "test_db"),
		},
	}
}

