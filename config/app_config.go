package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig menyimpan konfigurasi aplikasi
type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

// LoadAppConfig membaca variabel lingkungan dan mengembalikan instance AppConfig
func LoadAppConfig() AppConfig {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development" // Default environment
	}

	if appEnv != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	appName := os.Getenv("APP_NAME")

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000" // Default port
	}

	return AppConfig{
		AppName: appName,
		AppEnv:  appEnv,
		AppPort: appPort,
	}
}
