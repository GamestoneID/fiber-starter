package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"gamestone/services"

	"github.com/hibiken/asynq"
)

type GameTaskImpl struct {
	service services.GameService
}

// SyncGame implements GameTask.
func (g *GameTaskImpl) SyncGames(ctx context.Context, t *asynq.Task) error {
	fmt.Println("Syncing game...")

	if err := g.service.SyncGames(ctx); err != nil {
		fmt.Printf("failed to sync games: %v\n", err)
		return nil
	}

	fmt.Println("Syncing games completed successfully")

	return nil
}

// SyncGames implements GameTask.
func (g *GameTaskImpl) SyncGame(ctx context.Context, t *asynq.Task) error {
	var payload struct {
		Slug string `json:"slug"`
	}
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		fmt.Printf("failed to unmarshal payload: %v\n", err)
		return nil
	}

	if err := g.service.SyncGame(ctx, payload.Slug); err != nil {
		fmt.Printf("failed to sync game: %v\n", err)
		return nil
	}
	fmt.Println("Syncing game completed successfully")

	return nil

}

func NewGameTask(service services.GameService) GameTask {
	return &GameTaskImpl{service}
}
