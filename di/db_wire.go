//go:build wireinject
// +build wireinject

package di

import (
	"gamestone/database"

	"github.com/google/wire"
)

// Wire Set for Database
var DBSet = wire.NewSet(database.ConnectDB)
