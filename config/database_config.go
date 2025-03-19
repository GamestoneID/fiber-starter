package config

import (
	"os"
)

// Cloudflare menyimpan konfigurasi Cloudflare R2
type DatabaseConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBDriver   string
}

func LoadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBDriver:   os.Getenv("DB_DRIVER"),
	}
}
