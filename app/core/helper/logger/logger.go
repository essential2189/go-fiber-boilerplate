package logger

import (
	"fmt"
	"go-boilerplate/config"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Sugared struct {
	*zap.SugaredLogger
}

var Zap *Sugared

// Get initializes a zap.Logger Zap if it has not been initialized
// already and returns the same Zap for subsequent calls.
//func Get() *Sugared {
//	once.Do(func() {
//		Zap = New()
//	})
//
//	return Zap
//}

func New() *Sugared {
	conf, _ := config.NewConfig(config.NewVars())

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

	if os.Getenv("env") == "prd" {
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

	zapConfig := zap.Config{
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
			"type":    conf.App.Log.HistoryType,
			"version": conf.App.Version,
		},
	}

	Zap = &Sugared{
		zap.Must(zapConfig.Build()).Sugar(),
	}

	return Zap
}

func (l *Sugared) Printf(format string, args ...interface{}) {
	l.Infof(format, args...)
}
