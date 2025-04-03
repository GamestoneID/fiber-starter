//go:build wireinject
// +build wireinject

package di

import (
	"gamestone/app"

	"github.com/google/wire"
)

// Wire Set for App dependencies
var AppSet = wire.NewSet(
	app.NewFiberApp,
	ConfigSet,
	ProviderSet,
	HandlerSet,
)
