package helper

import (
	"go-boilerplate/app/core/helper/database"
	"go-boilerplate/app/core/helper/logger"
	"go-boilerplate/app/core/helper/resty"
	"go-boilerplate/app/core/helper/sentry"
	"go-boilerplate/app/core/helper/validator"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module(
	"helper",
	fx.Provide(
		logger.New,
		database.NewMaria,
		database.NewKafkaProducer,
		resty.NewRestyClient,
		sentry.New,
		validator.New,
	),
)

type Helper struct {
	fx.In

	MariaDB *gorm.DB
	Kafka   *database.Producer
	//Redis     *redis.Client
	Logger    *logger.Sugared
	Resty     *resty.HttpClient
	Validator *validator.Checker
}
