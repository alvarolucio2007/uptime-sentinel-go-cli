package testing

import (
	"testing"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/database"
	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
)

func TestPostgres(t *testing.T) {
	t.Run("Fluxo Completo Postgres", func(t *testing.T) {
	})
}

func testarPostgres(t *testing.T) {
	database.ConectarDatabase()
	entradaDB := [2]models.ModeloLink{
		{URL: "https://google.com", PeriodoSegundos: 10},
		{URL: "https://github.com", PeriodoSegundos: 10},
	}
	for _, entrada := range entradaDB {
	}
}
