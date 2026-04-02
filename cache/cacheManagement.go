// Package cache gerencia o Redis
package cache

import (
	"context"
	"encoding/json"
	"time"

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

func AdicionarLinkValkey(ID string, url *models.ModeloLink) error {
	dados, err := json.Marshal(url)
	if err != nil {
		models.ErroMarshalJSONSetCache.Log(err)
		return err
	}
	return RDB.Set(Ctx, ID, dados, 24*time.Hour).Err()
}

func BuscarLinkValkey(ID string) (string, error) {
	var resultado string
	var err error
	if resultado, err = RDB.Get(Ctx, ID).Result(); err != nil {
		models.ErroBuscarLinkCache.Log(err)
		return "", err
	}
	return resultado, nil
}

func DeletarLinkValkey(ID string) error {
	return RDB.Del(Ctx, ID).Err()
}
