package testing

import (
	"encoding/json"
	"strconv"
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
	urlOriginal := &models.ModeloLink{ID: 1000, URL: "https://google.com", PeriodoSegundos: 2, StatusEsperado: 200}
	IDString := strconv.Itoa(urlOriginal.ID)
	if err := cache.AdicionarLinkValkey(IDString, urlOriginal); err != nil {
		t.Errorf("Erro ao adicionar link ao valkey: %v", err)
	}
	jsonRecuperado, err := cache.BuscarLinkValkey(IDString)
	if err != nil {
		t.Errorf("Erro ao buscar link no valkey: %v", err)
	}
	var linkFinal models.ModeloLink
	if err := json.Unmarshal([]byte(jsonRecuperado), &linkFinal); err != nil {
		t.Errorf("Erro ao fazer Unmarshal do JSON recuperado: %v", err)
	}
	if linkFinal.URL != urlOriginal.URL {
		t.Errorf("Erro ao comparar URL recuperada e URL original, recebi %s e recebi %s", urlOriginal.URL, linkFinal.URL)
	}
	if err := cache.DeletarLinkValkey(IDString); err != nil {
		t.Errorf("Erro ao deletar link: %v", err)
	}
}
