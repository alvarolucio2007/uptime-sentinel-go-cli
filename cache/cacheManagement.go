// Package cache gerencia o Redis
package cache

import (
	"context"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	RDB *redis.Client
)

func ConectarCache() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	if err := RDB.Ping(Ctx).Err(); err != nil {
		models.ErroConexaoPingCache.Log(err)
		return err
	}
	return nil
}

func AdicionarLinkValkey(ID, url string) error {
	return RDB.Set(Ctx, ID, url, 0).Err()
}

func BuscarLinkRedis(ID string) (string, error) {
	return RDB.Get(Ctx, ID).Result()
}

func DeletarLinkRedis(ID string) error {
	return RDB.Del(Ctx, ID).Err()
}
