//go:build wireinject
// +build wireinject

package main

import (
	"gamestone/container"
	"gamestone/internal/di"

	"github.com/google/wire"
)

// Initialize App
func InitializeApp() (*container.AppContainer, error) {
	wire.Build(di.AppSet)
	return &container.AppContainer{}, nil
}
