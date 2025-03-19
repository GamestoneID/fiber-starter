package services

import (
	"gamestone/dto"
	"gamestone/utils"
)

type GameService interface {
	GetAllGames(filter *dto.FilterGame) (*utils.Pagination[dto.GameResponse], error)
	GetGameBySlug(slug string) (*dto.DetailGameResponse, error)
	CreateGame(game *dto.CreateGameRequest) (*dto.DetailGameResponse, error)
	UpdateGame(slug string, game *dto.UpdateGameRequest) (*dto.DetailGameResponse, error)
	DeleteGame(slug string) error
}
