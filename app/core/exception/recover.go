package exception

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"go-boilerplate/app/core/helper/sentry"
	"os"
	"runtime/debug"
)

func defaultStackTraceHandler(_ *fiber.Ctx, e interface{}) {
	_, _ = os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", e, debug.Stack())) //nolint:errcheck // This will never fail
}

// Helper function to set default values
func configDefault(config ...fiberRecover.Config) fiberRecover.Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return fiberRecover.ConfigDefault
	}

	// Override default config
	cfg := config[0]

	if cfg.EnableStackTrace && cfg.StackTraceHandler == nil {
		cfg.StackTraceHandler = defaultStackTraceHandler
	}

	return cfg
}

func Recover(config ...fiberRecover.Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	// Return new handler
	return func(c *fiber.Ctx) (err error) { //nolint:nonamedreturns // Uses recover() to overwrite the error
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Catch panics
		defer func() {
			if r := recover(); r != nil {
				if cfg.EnableStackTrace {
					cfg.StackTraceHandler(c, r)
				}

				var ok bool
				if err, ok = r.(error); ok {
					sentry.CaptureException(c, err)
				}
			}
		}()

		// Return err if exist, else move to next handler
		return c.Next()
	}
}
