//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package injector

import (
	controller "config-management-go/controller"
	"config-management-go/models/file"
	"config-management-go/models/iteration"
	fileService "config-management-go/service/file"
	iterationService "config-management-go/service/iteration"

	"github.com/google/wire"
)

func BuildApiInjector() (*ApiInjector, func(), error) {
	wire.Build(
		InitGormDB,
		iteration.NewRepository,
		file.NewRepository,
		fileService.NewService,
		iterationService.NewService,
		controller.NewIterationController,
		controller.NewFileController,
		NewControllerInjector,
		InitGinEngine,
		NewApiInjector,
	)
	return nil, nil, nil
}
