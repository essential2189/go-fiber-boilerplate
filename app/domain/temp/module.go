package temp

import (
	"go-boilerplate/app"
	"go-boilerplate/app/domain/temp/controller"

	"go.uber.org/fx"
)

var ControllerModule = fx.Module(
	"controller",
	fx.Provide(
		app.AsRoute(controller.NewController),
	),
)

var ServiceModule = fx.Module(
	"service",
	fx.Provide(
	//TempService,
	),
)

var RepositoryModule = fx.Module(
	"repository",
	fx.Provide(
	//TempRepository,
	),
)
