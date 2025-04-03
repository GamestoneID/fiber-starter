//go:build wireinject
// +build wireinject

package di

import (
	"gamestone/repositories"

	"github.com/google/wire"
)

// Wire Set
var RepositorySet = wire.NewSet(
	repositories.NewGameRepository,
)
