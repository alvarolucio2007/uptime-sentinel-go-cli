// Package exroutes faz as rotas externas do programa.
package exroutes

import "net/http"

func SetupExtRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /cadastrar")
}
