// Package verifiers verifica os links, periodo em segundos e status esperado.
package verifiers

import (
	"net/url"
	"slices"
	"strings"
)

func ValidarURL(URL string) bool {
	u, err := url.Parse(URL)
	if err != nil {
		return false
	}
	if u.Scheme == "" || u.Host == "" {
		return false
	}
	scheme := strings.ToLower(u.Scheme)
	return scheme == "http" || scheme == "https"
}

func ValidarStatus(statusChecar uint) bool {
	listaStatus := []uint{200, 201, 204, 301, 400, 401, 404}
	return slices.Contains(listaStatus, statusChecar)
}
