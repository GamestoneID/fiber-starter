package utils

import (
	"gamestone/dto"
	"gamestone/models"
)

func ToResponses[T any, R any](items []T, convertFunc func(*T) R) []R {
	if len(items) == 0 {
		return []R{} // Explicitly return an empty slice
	}

	responses := make([]R, 0, len(items))
	for i := range items {
		responses = append(responses, convertFunc(&items[i]))
	}
	return responses
}

func ToGameResponse(game *models.Game) dto.GameResponse {
	return dto.GameResponse{
		Slug:      game.Slug,
		Name:      game.Name,
		Image:     game.Image,
		IsActive:  game.IsActive,
		CreatedAt: game.CreatedAt,
		UpdatedAt: game.UpdatedAt,
	}
}

func ToDetailGameResponse(game *models.Game) dto.DetailGameResponse {
	return dto.DetailGameResponse{
		Slug:        game.Slug,
		Name:        game.Name,
		Description: game.Description,
		Image:       game.Image,
		IsActive:    game.IsActive,
		CreatedAt:   game.CreatedAt,
		UpdatedAt:   game.UpdatedAt,
	}
}
