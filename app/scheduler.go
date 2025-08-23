package app

import (
	"fmt"
	"gamestone/config"
	"gamestone/container"
	"log"

	"github.com/hibiken/asynq"
)

// NewAsynqScheduler creates a new asynq scheduler
func NewAsynqScheduler(
	appConfig *config.AppConfig,
	schedulerConfig *config.SchedulerConfig,
	redisConfig *config.RedisConfig,
	asynqClient *asynq.Client,
) (*container.SchedulerContainer, error) {
	redisClientOpt := asynq.RedisClientOpt{
		Addr: redisConfig.RedisHost + ":" + redisConfig.RedisPort,
	}
	defer asynqClient.Close()

	scheduler := asynq.NewScheduler(redisClientOpt, nil)

	scheduler.Register(
		schedulerConfig.SyncGamesCronExpression,
		asynq.NewTask("sync_games", nil),
	)

	log.Println("Task registered successfully")

	if err := scheduler.Start(); err != nil {
		return nil, fmt.Errorf("failed to start scheduler: %w", err)
	}

	return &container.SchedulerContainer{
		Scheduler: scheduler,
	}, nil
}
