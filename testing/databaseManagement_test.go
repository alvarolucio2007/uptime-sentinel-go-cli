package testing

import (
	"testing"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/database"
)

func TestPostgres(t *testing.T) {
	t.Run("Fluxo Completo Postgres", func(t *testing.T) {
	})
}

func testarPostgres(t *testing.T) {
	database.ConectarDatabase()
	entradaDB := [2]struct {
		link            string
		periodoSegundos uint
	}{
		{"https://google.com", 10},
		{"https://github.com", 10},
	}
	for _, entrada := range entradaDB {
	}
}
