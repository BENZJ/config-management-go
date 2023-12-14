//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package injector

import (
	controller "config-management-go/controller"
	"config-management-go/models/iteration"
	iterationService "config-management-go/service/iteration"

	"github.com/google/wire"
)

func BuildApiInjector() (*ApiInjector, func(), error) {
	wire.Build(
		InitGormDB,
		iteration.NewRepository,
		iterationService.NewService,
		controller.NewIterationController,
		NewControllerInjector,
		InitGinEngine,
		NewApiInjector,
	)
	return nil, nil, nil
}
