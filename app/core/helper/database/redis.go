package database

import (
	"context"
	"go-boilerplate/app/core/helper/logger"
	"go-boilerplate/config"

	"github.com/redis/go-redis/v9"
)

func NewRedis(config *config.Config) *redis.Client {
	ctx := context.Background()
	c := config.Infra.Redis
	conn := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    c.MasterName,
		SentinelAddrs: c.SentinelAddrs,
	})

	_, err := conn.Ping(ctx).Result()
	if err != nil {
		logger.Zap.Panicf("Failed to %v client connection. (%v)", c.MasterName, err)
	}

	/* For API Caching */
	//cacheClient := cache.New(&cache.Options{
	//	Redis: conn,
	//	Marshal: func(source interface{}) ([]byte, error) {
	//		return json.Marshal(source)
	//	},
	//	Unmarshal: func(source []byte, target interface{}) error {
	//		return json.Unmarshal(source, target)
	//	},
	//})

	return conn
}
