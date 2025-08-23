//go:build wireinject
// +build wireinject

package main

import (
	"gamestone/container"
	"gamestone/internal/di"

	"github.com/google/wire"
)

// InitializeScheduler builds WorkerContainer
func InitializeScheduler() (*container.SchedulerContainer, error) {
	wire.Build(di.SchedulerSet)
	return &container.SchedulerContainer{}, nil
}
