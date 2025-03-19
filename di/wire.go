//go:build wireinject
// +build wireinject

package di

import (
	"gamestone/container"

	"github.com/google/wire"
)

// Initialize App
func InitializeApp() (*container.AppContainer, error) {
	wire.Build(AppSet)
	return &container.AppContainer{}, nil
}
