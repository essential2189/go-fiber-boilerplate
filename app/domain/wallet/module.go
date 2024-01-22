package wallet

import (
	"go-boilerplate/app"
	"go-boilerplate/app/domain/wallet/controller"
	"go-boilerplate/app/domain/wallet/repository"
	"go-boilerplate/app/domain/wallet/service"
	"go.uber.org/fx"
)

var ControllerModule = fx.Module(
	"controller",
	fx.Provide(
		app.AsRoute(controller.NewWalletController),
	),
)

var ServiceModule = fx.Module(
	"service",
	fx.Provide(
		service.NewWalletService,
	),
)

var RepositoryModule = fx.Module(
	"repository",
	fx.Provide(
		repository.NewWalletRepository,
	),
)
