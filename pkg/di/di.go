package di

import (
	"framework-go/pkg/config"
	"framework-go/pkg/features/devOps"
	"framework-go/pkg/storage"

	"go.uber.org/dig"

	// the follows imports were injected for Hygen

	"framework-go/pkg/storage/ormo"
)

var container = dig.New()
var c *config.Config

func BuildContainer() *dig.Container {
	// config
	container.Provide(config.NewConfig)

	// the follows dialects were added for Hygen
	container.Provide(storage.NewDbOracle)

	// apiDevOps
	container.Provide(ormo.NewdevOpsRepo)
	container.Provide(devOps.NewdevOpsService)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
