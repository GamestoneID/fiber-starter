//go:build wireinject
// +build wireinject

package di

import (
	"gamestone/internal/tasks"

	"github.com/google/wire"
)

// Wire Set
var TaskSet = wire.NewSet(
	tasks.NewGameTask,
	ServiceSet,
)
