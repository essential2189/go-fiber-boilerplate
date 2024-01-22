package app

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/app/core/exception"
	"go-boilerplate/app/core/helper/logger"
	"go-boilerplate/config"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"
	"go.uber.org/zap/zapcore"
)

// NewFiber create a new Fiber application
func NewFiber(lc fx.Lifecycle, c *config.Config) *fiber.App {
	app := initializeFiber()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Zap.Infof("Server is starting on port: %s", c.Server.Port)

			addr := fmt.Sprintf(":%s", c.Server.Port)
			go app.Listen(addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}

func initializeFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:       "go-boilerplate",
		ServerHeader:  "go-boilerplate",
		Prefork:       false, // This will spawn multiple Go processes listening on the same port.
		CaseSensitive: true,  // When disabled, /Foo and /foo are treated the same.
		StrictRouting: true,  // When disabled, the router treats /foo and /foo/ as the same.
		UnescapePath:  true,  // url decoded path, ctx.Params(%key%)
		ErrorHandler:  exception.ErrorHandler,
	})
	app = setMiddleware(app)

	return app
}

func setMiddleware(app *fiber.App) *fiber.App {
	app.Use(recover.New())
	app.Use(cors.New())
	// Set request ID
	app.Use(requestid.New(requestid.Config{
		ContextKey: logger.RequestID,
	}))
	app.Use(logger.NewMiddleware(logger.Config{
		SkipURIs: []string{"/check_health"},
		Logger:   logger.Zap.Desugar(),
		Fields: []string{
			logger.RequestID,
			logger.Status,
			logger.Latency,
			logger.Error,
			logger.PID,
			logger.IP,
			logger.IPs,
			logger.Method,
			logger.URL,
			logger.Body,
			logger.Protocol,
			logger.ReqHeaders,
		},
		Levels: []zapcore.Level{zapcore.ErrorLevel, zapcore.ErrorLevel, zapcore.InfoLevel},
	}))
	app.Get("/check_health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	})
	app.Get("/swagger/*", swagger.HandlerDefault)

	return app
}
