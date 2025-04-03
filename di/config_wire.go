//go:build wireinject
// +build wireinject

package di

import (
	"gamestone/config"

	"github.com/google/wire"
)

// Wire Set for Database
var ConfigSet = wire.NewSet(
	config.LoadAppConfig,
	config.LoadDatabaseConfig,
	config.LoadRedisConfig,
)
