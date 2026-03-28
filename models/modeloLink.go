// Package models dá modelos fixos para melhor manutenção do sistema.
package models

type ModeloLink struct {
	ID              int    `json:"id"`
	URL             string `db:"URL" json:"url"`
	PeriodoSegundos uint   `db:"PeriodoSegundos" json:"periodo_segundos"`
	StatusEsperado  int    `json:"status_esperado"`
}
