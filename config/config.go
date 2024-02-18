package config

import (
	"fmt"
	"os"
)

// Config holds the application configuration
type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

// NewConfig creates a new Config instance with default values
func NewConfig() *Config {
	return &Config{
		DBUsername: getEnv("DB_USERNAME", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "translation_db"),
	}
}

// getEnv retrieves the value of an environment variable or returns a default value if not set
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// DatabaseURL constructs the database connection URL
func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUsername, c.DBPassword, c.DBName)
}
