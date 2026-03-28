// Package cache gerencia o Redis
package cache

import (
	"context"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	RDB *redis.Client
)

func ConectarCache() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	if err := RDB.Ping(ctx).Err(); err != nil {
		models.ErroConexaoPingCache.Log(err)
		return err
	}
	return nil
}
