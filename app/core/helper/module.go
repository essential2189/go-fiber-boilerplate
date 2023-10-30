package helper

import (
	"go-boilerplate/app/core/helper/database"
	"go-boilerplate/app/core/helper/resty"
	"go-boilerplate/app/core/helper/sentry"
	"go-boilerplate/app/core/helper/validator"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"helper",
	fx.Provide(
		database.NewMaria,
		database.NewKafkaProducer,
		database.NewRedis,
		validator.New,
		sentry.New,
		resty.NewRestyClient,
	),
)
