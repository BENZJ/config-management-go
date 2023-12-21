//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package injector

import (
	controller "config-management-go/controller"
	"config-management-go/models/file"
	fileitem "config-management-go/models/file_item"
	"config-management-go/models/iteration"
	fileService "config-management-go/service/file"
	fileItemService "config-management-go/service/file_item"
	iterationService "config-management-go/service/iteration"

	"github.com/google/wire"
)

func BuildApiInjector() (*ApiInjector, func(), error) {
	wire.Build(
		InitGormDB,
		iteration.NewRepository,
		file.NewRepository,
		fileitem.NewRepository,
		fileService.NewService,
		iterationService.NewService,
		fileItemService.NewService,
		controller.NewIterationController,
		controller.NewFileController,
		controller.NewFileItemController,
		NewControllerInjector,
		InitGinEngine,
		NewApiInjector,
	)
	return nil, nil, nil
}
