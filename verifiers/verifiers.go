// Package verifiers verifica os links, periodo em segundos e status esperado.
package verifiers

import (
	"errors"
	"net/url"
	"slices"
	"strings"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
)

func validarURL(URL string) error {
	u, err := url.Parse(URL)
	if err != nil {
		return err
	}
	if u.Scheme == "" || u.Host == "" {
		return errors.New("erro ao verificar se o host ou scheme do URL")
	}
	scheme := strings.ToLower(u.Scheme)
	if scheme == "http" || scheme == "https" {
		return nil
	}
	return errors.New("scheme não é http nem https ")
}

func validarStatus(statusChecar uint) error {
	listaStatus := []uint{200, 201, 204, 301, 400, 401, 404}
	if !slices.Contains(listaStatus, statusChecar) {
		return errors.New("status não está na lista de status aceitos")
	}
	return nil
}

func ValidacaoCompleta(linkJSON *models.ModeloLink) error {
	if err := validarURL(linkJSON.URL); err != nil {
		models.ErroValidacaoURL.Log(err)
		return err
	}
	if err := validarStatus(linkJSON.StatusEsperado); err != nil {
		models.ErroValidacaoStatus.Log(err)
		return err
	}
	return nil
}
