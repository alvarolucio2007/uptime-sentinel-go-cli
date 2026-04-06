// Package models dá modelos fixos para melhor manutenção do sistema.
package models

type ModeloLink struct {
	ID              int    `gorm:"primaryKey" json:"id"`
	URL             string `gorm:"URL" json:"url"`
	PeriodoSegundos uint   `gorm:"PeriodoSegundos" json:"periodo_segundos"`
	StatusEsperado  uint   `json:"status_esperado"`
}
