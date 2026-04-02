package inroutes

import (
	"encoding/json"
	"net/http"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
)

func HandlerAdd(w http.ResponseWriter, r *http.Request) error {
	var input models.ModeloLink
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, models.ErroDecodeJSONAPIExterna.Mensagem, http.StatusBadRequest)
		return err
	}
}
