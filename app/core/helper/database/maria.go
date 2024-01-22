package database

import (
	"fmt"
	"go-boilerplate/app/core/helper/logger"
	"go-boilerplate/config"
	"net/url"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/samber/lo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

// NewMaria opens a database connection.
func NewMaria(config *config.Config) *gorm.DB {
	gormConfig := newGormConfig()
	sources, replicas := getDialectors(config.Infra.DB)

	db, err := gorm.Open(sources[0], gormConfig)
	if err != nil {
		logger.Zap.Fatalf("failed to connect to database : %+v", err)
	}
	err = db.Use(getDBResolver(sources, replicas))
	if err != nil {
		logger.Zap.Fatalf("failed to set dbresolver : %+v", err)
	}

	err = checkConnection(db)
	if err != nil {
		logger.Zap.Fatalf("failed to check connection: %+v", err)
	}

	logger.Zap.Info("DB connected")

	return db
}

func getDBResolver(sources []gorm.Dialector, replicas []gorm.Dialector) *dbresolver.DBResolver {
	return dbresolver.
		Register(dbresolver.Config{
			Sources:  sources,
			Replicas: replicas,
		}).
		SetMaxIdleConns(10).
		SetMaxOpenConns(50).
		SetConnMaxLifetime(59 * time.Second)
}

// checkConnection checks if the connection is alive.
func checkConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return errors.Wrap(err, "failed to get *sql.DB")
	}
	if err := sqlDB.Ping(); err != nil {
		return errors.Wrap(err, "failed to ping")
	}

	return nil
}

// getDialectors returns sources and replicas dialectors.
func getDialectors(dbConfig config.DB) ([]gorm.Dialector, []gorm.Dialector) {
	iteratee := func(addr string, index int) gorm.Dialector {
		dsn := sprintDSN(
			dbConfig.User,
			dbConfig.Password,
			addr,
			dbConfig.DBName,
		)
		return mysql.Open(dsn)
	}
	sources := lo.Map(dbConfig.SourceAddrs, iteratee)
	replicas := lo.Map(dbConfig.ReplicaAddrs, iteratee)

	return sources, replicas
}

// newGormConfig returns a new gorm.Config.
func newGormConfig() *gorm.Config {
	return &gorm.Config{Logger: newLogger()}
}

// newLogger returns a new gorm/logger.Interface.
func newLogger() gormLogger.Interface {
	return gormLogger.New(
		logger.Zap,
		gormLogger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  getLogLevel(), // Log level
			IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,         // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)
}

// sprintDSN returns a DSN(data source name) string.
func sprintDSN(user string, pass string, address string, dbname string) string {
	values := url.Values{}
	values.Add("charset", "utf8mb4")
	values.Add("parseTime", "true")
	query := values.Encode()

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", user, pass, address, dbname, query)
}

func getLogLevel() gormLogger.LogLevel {
	var level gormLogger.LogLevel

	switch config.Profile {
	case "dev":
		level = gormLogger.Info
	case "local":
		level = gormLogger.Info
	default:
		level = gormLogger.Error
	}

	return level
}
