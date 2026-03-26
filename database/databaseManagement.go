// Package database conecta à base de dados
package database

import (
	"database/sql"
	"fmt"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConectarDatabase() error {
	var err error
	dsn := "host=pg_sentinel user=user password=password dbname=sentinel sslmode=disabled"
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		models.ErroAberturaPostgres.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		models.ErroConexaoPostgres.Fatal(err)
	}
	fmt.Println("Postgres conectado")
	return nil
}

func MigrarBanco() error {
	query := `CREATE TABLE IF NOT EXISTS linksSentinel (
		URL TEXT UNIQUE NOT NULL,
		PeriodoSegundos INTEGER NOT NULL DEFAULT 0
	);`
	if _, err := DB.Exec(query); err != nil {
	}
}
