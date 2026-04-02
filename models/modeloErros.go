package models

import "log"

type GerenciadorErro struct {
	Mensagem string
}

func (e GerenciadorErro) Log(err error) {
	log.Printf("⚠️ %s | Detalhes: %v", e.Mensagem, err)
}

func (e GerenciadorErro) Fatal(err error) {
	log.Fatalf("❌ FATAL: %s | Detalhes: %v", e.Mensagem, err)
}

var (
	// Erros Setup e migração do PostgreSQL
	ErroAberturaPostgres = GerenciadorErro{"Erro ao tentar abrir o Postgres: "}
	ErroConexaoPostgres  = GerenciadorErro{"Erro ao tentar conectar ao Postgres: "}
	ErroMigracaoPostgres = GerenciadorErro{"Erro na migração do Postgres: "}

	// Erros de inserção no PostgreSQL
	ErroEntradaPostgres = GerenciadorErro{"Erro ao criar item: "}
	ErroDeletePostgres  = GerenciadorErro{"Erro ao deletar item: "}
)

// Erros cache
var (
	ErroAberturaCache       = GerenciadorErro{"Erro ao tentar abrir o Cache: "}
	ErroConexaoPingCache    = GerenciadorErro{"Erro ao receber o ping do cache: "}
	ErroMarshalJSONSetCache = GerenciadorErro{"Erro ao transformar dados em JSON: "}
	ErroSetCache            = GerenciadorErro{"Erro ao dar set no cache: "}
	ErroGetCache            = GerenciadorErro{"Erro ao pegar a informação do cache: "}
	ErroBuscarLinkCache     = GerenciadorErro{"Erro ao buscar link no cache: "}
	ErroDeletarLinkCache    = GerenciadorErro{"Erro ao deletar link no cache: "}
)

// Erros API Externa
var (
	ErroDecodeJSONAPIExterna = GerenciadorErro{"Erro ao tentar decodar JSON da API externa: "}
)

// Erros validação
var (
	ErroValidacaoURL = GerenciadorErro{"Erro ao validar o URL: "}
)
