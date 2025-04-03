//go:build wireinject
// +build wireinject

package di

import (
	"gamestone/providers"

	"github.com/google/wire"
)

// Wire Set for Provider
var ProviderSet = wire.NewSet(
	providers.NewDBConnection,
	providers.NewAsyncClient,
	providers.NewValidator,
)
