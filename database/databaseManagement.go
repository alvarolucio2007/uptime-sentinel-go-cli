// Package database conecta à base de dados
package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConectarDatabase() {
	var err error
	dsn := "host=pg_sentinel user=user password=password dbname=sentinel sslmode=disabled"
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Não foi possível se conectar ao postgres: ", err)
	}
	fmt.Println("Postgres conectado")
}
