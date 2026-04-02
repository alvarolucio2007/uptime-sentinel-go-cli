package exroutes

import (
	"encoding/json"
	"net/http"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
	"github.com/alvarolucio2007/uptime-sentinel-go-cli/verifiers"
)

func HandlerAdd(w http.ResponseWriter, r *http.Request) error {
	var input *models.ModeloLink
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, models.ErroDecodeJSONAPIExterna.Mensagem, http.StatusBadRequest)
		return err
	}
	if err := verifiers.ValidacaoCompleta(input); err != nil {
		http.Error(w, models.ErroValidacaoStatus.Mensagem, http.StatusBadRequest)
		return err
	}
}
