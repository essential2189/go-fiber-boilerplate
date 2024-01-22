package temp

import (
	"go-boilerplate/app"
	"go-boilerplate/app/domain/temp/controller"
	"go-boilerplate/app/domain/temp/service"

	"go.uber.org/fx"
)

var ControllerModule = fx.Module(
	"controller",
	fx.Provide(
		app.AsRoute(controller.NewTempController),
	),
)

var ServiceModule = fx.Module(
	"service",
	fx.Provide(
		service.NewTempService,
	),
)
