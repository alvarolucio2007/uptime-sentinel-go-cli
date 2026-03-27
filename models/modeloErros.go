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
)
