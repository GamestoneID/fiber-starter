package dto

import (
	"time"
)

type CreateGameRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	Image       string  `json:"image" validate:"required"`
	IsActive    bool    `json:"is_active"`
}

type UpdateGameRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	IsActive    *bool   `json:"is_active"`
}

type FilterGame struct {
	IsActive *bool `json:"is_active"`
	Page     *uint `json:"page"`
	Limit    *uint `json:"limit"`
}

type GameResponse struct {
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DetailGameResponse struct {
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Image       string    `json:"image"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
