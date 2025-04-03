package tasks

import (
	"context"

	"github.com/hibiken/asynq"
)

type GameTask interface {
	SyncGames(ctx context.Context, t *asynq.Task) error
	SyncGame(ctx context.Context, t *asynq.Task) error
}
