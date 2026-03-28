// Package database conecta à base de dados
package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConectarDatabase() error {
	var err error
	dsn := "host=localhost user=user password=password dbname=sentinel sslmode=disable"
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
		ID SERIAL PRIMARY KEY,
		URL TEXT UNIQUE NOT NULL,
		PeriodoSegundos INTEGER NOT NULL DEFAULT 0,
		StatusEsperado INTEGER NOT NULL
	);`
	if _, err := DB.Exec(query); err != nil {
		return err
	}
	fmt.Println("Tabela pronta para uso.")
	return nil
}

func CriarEntradaPostgres(URL string, periodo, statusEsperado uint) error {
	var id uint
	query := "INSERT INTO linksSentinel (URL,PeriodoSegundos,StatusEsperado) VALUES ($1,$2,$3) RETURNING id"
	if err := DB.QueryRow(query, URL, periodo, statusEsperado).Scan(&id); err != nil {
		models.ErroEntradaPostgres.Log(err)
		return err
	}
	return nil
}

func DeletarEntradaPostgres(ID string) error {
	query := "DELETE FROM linksSentinel WHERE ID=$1"
	res, err := DB.Exec(query, ID)
	if err != nil {
		models.ErroDeletePostgres.Log(err)
		return err
	}
	qtd, _ := res.RowsAffected()
	if qtd == 0 {
		log.Println("ID inexistente, nada deletado.")
	}
	return nil
}
