//go:build wireinject
// +build wireinject

package main

import (
	"gamestone/container"
	"gamestone/internal/di"

	"github.com/google/wire"
)

// Initialize Worker
func InitializeWorker() (*container.WorkerContainer, error) {
	wire.Build(di.WorkerSet)
	return &container.WorkerContainer{}, nil
}
