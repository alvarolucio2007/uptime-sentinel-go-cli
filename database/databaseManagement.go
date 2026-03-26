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
	dsn := "host=pg_sentinel user=user password=password dbname=sentinel sslmode=disabled"
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Printf("%s. Detalhes: %+v ", models.ErrosSlice["ErroAberturaPostgres"], err)
	}
	err = DB.Ping()
	if err != nil {
		log.Printf("%s. Detalhes:")
	}
	fmt.Println("Postgres conectado")
	return nil
}
