package config

import (
	"os"
)

// DatabaseConfig menyimpan konfigurasi database
type DatabaseConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBDriver   string
}

// LoadDatabaseConfig membaca konfigurasi dari environment variable
func LoadDatabaseConfig() (*DatabaseConfig, error) {
	return &DatabaseConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBDriver:   os.Getenv("DB_DRIVER"),
	}, nil
}
