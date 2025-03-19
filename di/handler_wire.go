//go:build wireinject
// +build wireinject

package di

import (
	"gamestone/handlers"

	"github.com/google/wire"
)

// Wire Set
var HandlerSet = wire.NewSet(
	handlers.NewGameHandler,
	ServiceSet,
	ValidatorSet,
)
