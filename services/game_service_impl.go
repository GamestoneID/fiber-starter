package services

import (
	"gamestone/dto"
	"gamestone/models"
	"gamestone/repositories"
	"gamestone/utils"
)

type GameServiceImpl struct {
	repository repositories.GameRepository
}

// CreateGame implements GameService.
func (g *GameServiceImpl) CreateGame(game *dto.CreateGameRequest) (*dto.DetailGameResponse, error) {
	gameToCreate := &models.Game{
		Slug:        utils.GenerateSlug(game.Name),
		Name:        game.Name,
		Description: game.Description,
		Image:       game.Image,
		IsActive:    game.IsActive,
	}

	createdGame, err := g.repository.Create(gameToCreate)
	if err != nil {
		return nil, err
	}

	response := utils.ToDetailGameResponse(createdGame)
	return &response, nil
}

// DeleteGame implements GameService.
func (g *GameServiceImpl) DeleteGame(slug string) error {
	return g.repository.Delete(slug)
}

// GetAllGames implements GameService.
func (g *GameServiceImpl) GetAllGames(filter *dto.FilterGame) (*utils.Pagination[dto.GameResponse], error) {
	games, metadata, err := g.repository.FindAll(filter)
	if err != nil {
		return nil, err
	}

	data := utils.ToResponses(games, utils.ToGameResponse)

	response := &utils.Pagination[dto.GameResponse]{
		Data:     data,
		Metadata: *metadata,
	}

	return response, nil
}

// GetGameBySlug implements GameService.
func (g *GameServiceImpl) GetGameBySlug(slug string) (*dto.DetailGameResponse, error) {
	game, err := g.repository.FindBySlugWithVariantsAndProducts(slug)
	if err != nil {
		return nil, err
	}

	response := utils.ToDetailGameResponse(game)
	return &response, nil
}

// UpdateGame implements GameService.
func (g *GameServiceImpl) UpdateGame(slug string, game *dto.UpdateGameRequest) (*dto.DetailGameResponse, error) {
	gameToUpdate, err := g.repository.FindBySlug(slug)

	if err != nil {
		return nil, err
	}

	if game.Name != nil {
		gameToUpdate.Name = *game.Name
	}

	if game.Description != nil {
		gameToUpdate.Description = game.Description
	}

	if game.Image != nil {
		gameToUpdate.Image = *game.Image
	}

	if game.IsActive != nil {
		gameToUpdate.IsActive = *game.IsActive
	}

	gameUpdated, err := g.repository.Update(gameToUpdate)
	if err != nil {
		return nil, err
	}

	response := utils.ToDetailGameResponse(gameUpdated)
	return &response, nil
}

func NewGameService(repository repositories.GameRepository) GameService {
	return &GameServiceImpl{repository}
}
