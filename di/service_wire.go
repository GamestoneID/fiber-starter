//go:build wireinject
// +build wireinject

package di

import (
	"gamestone/services"

	"github.com/google/wire"
)

// Wire Set
var ServiceSet = wire.NewSet(
	services.NewGameService,
	RepositorySet,
)
