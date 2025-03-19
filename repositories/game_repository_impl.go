package repositories

import (
	"gamestone/dto"
	"gamestone/models"
	"gamestone/utils"
	"math"

	"gorm.io/gorm"
)

type GameRepositoryImpl struct {
	db *gorm.DB
}

// FindBySlugWithVariants implements GameRepository.
func (g GameRepositoryImpl) FindBySlugWithVariantsAndProducts(slug string) (*models.Game, error) {
	var game models.Game

	if err := g.db.Preload("Variants.Product").Where("slug = ?", slug).First(&game).Error; err != nil {
		return nil, err
	}

	return &game, nil
}

// FindByVocaProductId implements GameRepository.
func (g GameRepositoryImpl) FindByVocaProductId(vocaProductId uint) (*models.Game, error) {
	var game models.Game

	if err := g.db.Where("voca_product_id = ?", vocaProductId).First(&game).Error; err != nil {
		return nil, err
	}

	return &game, nil
}

// Create implements GameRepository.
func (g GameRepositoryImpl) Create(game *models.Game) (*models.Game, error) {
	if err := g.db.Create(game).Error; err != nil {
		return nil, err
	}

	return game, nil
}

// Delete implements GameRepository.
func (g GameRepositoryImpl) Delete(slug string) error {
	err := g.db.Where("slug = ?", slug).Delete(&models.Game{}).Error

	return err
}

// FindAll implements GameRepository.
func (g GameRepositoryImpl) FindAll(filter *dto.FilterGame) ([]models.Game, *utils.MetadataPagination, error) {
	var products []models.Game
	var total int64

	query := g.db.Model(&models.Game{})

	if filter.IsActive != nil {
		query = query.Where("is_active = ?", *filter.IsActive)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}

	limit, page, offset := utils.GetPagination(filter.Page, filter.Limit)
	query = query.Limit(limit).Offset(offset)

	if err := query.Find(&products).Error; err != nil {
		return nil, nil, err
	}

	metadataPagination := &utils.MetadataPagination{
		Total:      uint(total),
		Page:       uint(page),
		Limit:      uint(limit),
		TotalPages: uint(math.Ceil(float64(total) / float64(limit))),
	}

	return products, metadataPagination, nil
}

// FindBySlug implements GameRepository.
func (g GameRepositoryImpl) FindBySlug(slug string) (*models.Game, error) {
	var game models.Game

	if err := g.db.Where("slug = ?", slug).First(&game).Error; err != nil {
		return nil, err
	}

	return &game, nil
}

// Update implements GameRepository.
func (g GameRepositoryImpl) Update(game *models.Game) (*models.Game, error) {
	if err := g.db.Save(game).Error; err != nil {
		return nil, err
	}

	return game, nil
}

func NewGameRepository(db *gorm.DB) GameRepository {
	return GameRepositoryImpl{db}
}
