package config

import (
	"os"
)

// Cloudflare menyimpan konfigurasi Cloudflare R2
type RedisConfig struct {
	RedisHost string
	RedisPort string
}

func LoadRedisConfig() RedisConfig {
	return RedisConfig{
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
	}
}
