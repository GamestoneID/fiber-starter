package app

import (
	"gamestone/config"
	"gamestone/container"
	"gamestone/tasks"

	"github.com/hibiken/asynq"
)

// NewAsynqWorker initializes the Asynq worker with all necessary routes
func NewAsynqWorker(
	appConfig config.AppConfig,
	redisConfig config.RedisConfig,
	gameTask tasks.GameTask,
) (*container.WorkerContainer, error) {
	redisClientOpt := asynq.RedisClientOpt{
		Addr: redisConfig.RedisHost + ":" + redisConfig.RedisPort,
	}
	srv := asynq.NewServer(redisClientOpt, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc("sync_games", gameTask.SyncGames)
	mux.HandleFunc("sync_game", gameTask.SyncGame)

	return &container.WorkerContainer{
		Server: srv,
		Mux:    mux,
	}, nil
}
