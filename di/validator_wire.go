//go:build wireinject
// +build wireinject

package di

import (
	"gamestone/validator"

	"github.com/google/wire"
)

// Wire Set for Database
var ValidatorSet = wire.NewSet(validator.ProvideValidator)
