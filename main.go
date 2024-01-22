package main

import (
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/app"
	"go-boilerplate/app/core"
	"go-boilerplate/app/core/helper"
	"go-boilerplate/app/core/helper/logger"
	"go-boilerplate/app/domain/temp"
	"go-boilerplate/app/domain/wallet"
	"go-boilerplate/config"
	"go.uber.org/fx"
)

// @title go-boilerplate API
// @version 1.0
// @host localhost:8080
// @accept application/json
// @produce application/json
func main() {
	fx.New(
		config.Module,
		helper.Module,

		core.RepositoryModule,
		core.ClientModule,
		core.BaseModule,

		temp.ControllerModule,
		temp.ServiceModule,

		wallet.ControllerModule,
		wallet.ServiceModule,
		wallet.RepositoryModule,

		fx.Provide(
			app.NewFiber,
			fx.Annotate(
				app.NewRouter,
				fx.ParamTags(``, `group:"routes"`),
			),
		),
		fx.Invoke(
			func(*logger.Sugared) {},
			func(fiber.Router) {},
			func(*fiber.App) {},
		),
	).Run()
}
