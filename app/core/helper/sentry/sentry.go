package sentry

import (
	"fmt"
	"go-boilerplate/app/core/consts"
	"go-boilerplate/config"
	"log"
	"runtime"
	"time"

	_sentry "github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
)

func CaptureException(ctx *fiber.Ctx, err error) {
	capture(ctx, err)
}

func New(config *config.Config) fiber.Handler {
	dsn := config.Sentry.Dsn

	log.Println("sentry initializing... dsn is : ", dsn)
	err := _sentry.Init(getOptions(dsn))
	if err != nil {
		log.Fatal(fmt.Errorf("sentry init error : %v", err))
		return nil
	}

	return func(ctx *fiber.Ctx) error {
		return ctx.Next()
	}
}

func capture(ctx *fiber.Ctx, err error) {
	stackTrace := ""
	stackBuffer := make([]byte, 4096)
	stackSize := runtime.Stack(stackBuffer, false)
	stackTrace = string(stackBuffer[:stackSize])

	if config.Profile != consts.EnvDev || config.Test {
		setScopeAndCapture(ctx, err, stackTrace)
		_sentry.Flush(time.Second * 4)
	}
}
