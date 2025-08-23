package di

import (
	"gamestone/app"

	"github.com/google/wire"
)

// Wire Set for Worker dependencies
var SchedulerSet = wire.NewSet(
	app.NewAsynqScheduler,
	ConfigSet,
	ProviderSet,
)
