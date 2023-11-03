package main

import (
	"go-boilerplate/app"
	"go-boilerplate/app/core/helper"
	"go-boilerplate/app/core/helper/logger"
	"go-boilerplate/app/domain/temp"
	"go-boilerplate/config"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// @title go-boilerplate API
// @version 1.0
// @contact.name 프로덕트 도메인
// @contact.email product.domain.tech@wavve.com
// @host localhost:8080
// @accept application/json
// @produce application/json
func main() {
	fx.New(
		temp.ControllerModule,
		temp.ServiceModule,
		temp.RepositoryModule,

		config.Module,
		helper.Module,
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
