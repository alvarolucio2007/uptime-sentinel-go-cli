package testing

import (
	"testing"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/cache"
)

func TestValkey(t *testing.T) {
	t.Run("Fluxo Completo Valkey", func(t *testing.T) {
		// colocar função aqui...
	})
}

func testarValkey(t *testing.T) {
	if err := cache.ConectarCache(); err != nil {
		t.Errorf("Erro ao se concectar ao cache: %v", err)
	}
}
