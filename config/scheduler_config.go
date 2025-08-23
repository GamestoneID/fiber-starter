package config

import (
	"os"
)

// Cloudflare menyimpan konfigurasi Cloudflare R2
type SchedulerConfig struct {
	SyncGamesCronExpression string
}

func LoadSchedulerConfig() (*SchedulerConfig, error) {
	return &SchedulerConfig{
		SyncGamesCronExpression: os.Getenv("SYNC_GAMES_CRON_EXPRESSION"),
	}, nil
}
