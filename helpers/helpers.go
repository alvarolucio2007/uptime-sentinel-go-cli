// Package helpers tem funções universais de auxílio
package helpers

import (
	"log"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
)

func GenerateErrorLog(errorName string, err error) {
	log.Printf("%s. Detalhes: %+v ", models.ErrosSlice[errorName], err)
}
