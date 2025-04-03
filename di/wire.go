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

// Initialize Worker
func InitializeWorker() (*container.WorkerContainer, error) {
	wire.Build(WorkerSet)
	return &container.WorkerContainer{}, nil
}
