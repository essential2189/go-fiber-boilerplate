package logger

import (
	"fmt"
	"go-boilerplate/config"
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	historyType = "history-store"
	version     = "1.0.0"
)

var once sync.Once

type Sugared struct {
	*zap.SugaredLogger
}

var instance *Sugared

// Get initializes a zap.Logger instance if it has not been initialized
// already and returns the same instance for subsequent calls.
func Get() *Sugared {
	once.Do(func() {
		instance = New()
	})

	return instance
}

func New() *Sugared {
	level := zap.InfoLevel
	if config.LogLevel != "" {
		levelFromEnv, err := zapcore.ParseLevel(config.LogLevel)
		if err != nil {
			log.Println(
				fmt.Errorf("invalid level, defaulting to INFO: %w", err),
			)
		}

		level = levelFromEnv
	}

	var encoderConfig zapcore.EncoderConfig
	var encoding string

	if os.Getenv("wavve_env") == "prd" {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoding = "json"
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoding = "console"
	}

	encoderConfig.MessageKey = "message"
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.CallerKey = "caller"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          encoding,
		EncoderConfig:     encoderConfig,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"type":    historyType,
			"version": version,
		},
	}

	instance = &Sugared{
		zap.Must(config.Build()).Sugar(),
	}

	return instance
}

func (l *Sugared) Printf(format string, args ...interface{}) {
	l.Infof(format, args...)
}
