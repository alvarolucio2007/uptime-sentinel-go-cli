package testing

import (
	"testing"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
	"github.com/alvarolucio2007/uptime-sentinel-go-cli/verifiers"
)

func TestVerifiers(t *testing.T) {
	t.Run("Fluxo completo de testes", func(t *testing.T) {
		testVerificarLink(t)
	})
}

func testVerificarLink(t *testing.T) {
	linksTeste := []struct {
		nome       string
		modeloLink models.ModeloLink
		passa      bool
	}{
		{"Passa", models.ModeloLink{URL: "https://google.com", PeriodoSegundos: 2, StatusEsperado: 200}, true},
		{"Não passa: sem http nem https", models.ModeloLink{URL: "google.com", PeriodoSegundos: 2, StatusEsperado: 200}, false},
		{"Não passa: sem host nem scheme", models.ModeloLink{URL: "teste", PeriodoSegundos: 2, StatusEsperado: 200}, false},
		{"Não passa: status inválido", models.ModeloLink{URL: "https://google.com", PeriodoSegundos: 2, StatusEsperado: 125}, false},
	}
	for _, testCase := range linksTeste {
		err := (verifiers.ValidacaoCompleta(&testCase.modeloLink) != nil)
		if err == testCase.passa {
			t.Errorf("Erro: resultado de %s diferente do esperado %t", testCase.modeloLink.URL, testCase.passa)
		}
	}
}
