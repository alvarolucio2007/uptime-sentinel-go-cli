package testing

import (
	"testing"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/cache"
	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
)

func TestValkey(t *testing.T) {
	t.Run("Fluxo Completo Valkey", func(t *testing.T) {
		testarValkey(t)
	})
}

func testarValkey(t *testing.T) {
	if err := cache.ConectarCache(); err != nil {
		t.Errorf("Erro ao se concectar ao cache: %v", err)
	}
	ID, urlOriginal := "1000", &models.ModeloLink{ID: 1000, URL: "https://google.com", PeriodoSegundos: 2, StatusEsperado: 200}
	if err := cache.AdicionarLinkValkey(ID, urlOriginal); err != nil {
		t.Errorf("Erro ao adicionar link ao valkey: %v", err)
	}
	urlRecuperada, err := cache.BuscarLinkValkey("1000")
	if err != nil {
		t.Errorf("Erro ao buscar link no valkey: %v", err)
	}
	if urlRecuperada != urlOriginal.URL {
		t.Errorf("Erro ao comparar URL recuperada e URL original, recebi %s e recebi %s", urlOriginal.URL, urlRecuperada)
	}
}
