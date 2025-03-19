package repositories

import (
	"gamestone/dto"
	"gamestone/models"
	"gamestone/utils"
)

type GameRepository interface {
	FindAll(filter *dto.FilterGame) ([]models.Game, *utils.MetadataPagination, error)
	FindByVocaProductId(vocaProductId uint) (*models.Game, error)
	FindBySlug(slug string) (*models.Game, error)
	FindBySlugWithVariantsAndProducts(slug string) (*models.Game, error)
	Create(game *models.Game) (*models.Game, error)
	Update(game *models.Game) (*models.Game, error)
	Delete(slug string) error
}
