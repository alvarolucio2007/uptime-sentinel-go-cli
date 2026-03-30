package testing

import (
	"testing"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/database"
	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
)

func TestPostgres(t *testing.T) {
	t.Run("Fluxo Completo Postgres", func(t *testing.T) {
		testarPostgres(t)
	})
}

func testarPostgres(t *testing.T) {
	if err := database.ConectarDatabase(); err != nil {
		t.Errorf("Erro ao se conectar à base de dados: %v", err)
	}
	query := "DROP TABLE IF EXISTS linksSentinel"
	if _, err := database.DB.Exec(query); err != nil {
		t.Errorf("Erro ao limpar a base de dados: %v", err)
	}
	if err := database.MigrarBanco(); err != nil {
		t.Errorf("Erro ao migrar a base de dados: %v", err)
	}
	entradaDB := [2]models.ModeloLink{
		{URL: "https://google.com", PeriodoSegundos: 10, StatusEsperado: uint(101)}, // StatusEsperado só placeholder msm kk
		{URL: "https://github.com", PeriodoSegundos: 10, StatusEsperado: uint(101)},
	}
	for _, entrada := range entradaDB {
		if err := database.CriarEntradaPostgres(entrada.URL, entrada.PeriodoSegundos, entrada.StatusEsperado); err != nil {
			t.Errorf("Erro ao adicionar livro: %v", err)
		}
	}
}
