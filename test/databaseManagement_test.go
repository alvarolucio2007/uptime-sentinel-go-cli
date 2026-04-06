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
		t.Fatalf("Erro ao se conectar à base de dados: %v", err)
	}
	if err := database.DB.Migrator().DropTable(&models.ModeloLink{}); err != nil {
		t.Fatalf("Erro ao dar drop table: %v", err)
	}
	if err := database.DB.AutoMigrate(&models.ModeloLink{}); err != nil {
		t.Errorf("Erro ao migrar database: %v", err)
	}
	entradaDB := [2]models.ModeloLink{
		{URL: "https://google.com", PeriodoSegundos: 10, StatusEsperado: uint(101)}, // StatusEsperado só placeholder msm kk
		{URL: "https://github.com", PeriodoSegundos: 10, StatusEsperado: uint(101)},
	}
	for _, e := range entradaDB {
		if err := database.CriarEntradaPostgres(e); err != nil {
			t.Errorf("Falha ao inserir %s:%v", e.URL, err)
		}
	}
}
