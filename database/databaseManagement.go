// Package database conecta à base de dados
package database

import (
	"fmt"
	"os"

	"github.com/alvarolucio2007/uptime-sentinel-go-cli/models"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarDatabase() error {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=user password=password dbname=sentinel sslmode=disable"
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		models.ErroAberturaPostgres.Fatal(err)
	}
	if err = DB.AutoMigrate(&models.ModeloLink{}); err != nil {
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

func CriarEntradaPostgres(dados models.ModeloLink) error {
	if resultado := DB.Create(&dados); resultado.Error != nil {
		models.ErroEntradaPostgres.Log(resultado.Error)
		return resultado.Error
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
		fmt.Println("ID inexistente, nada deletado.")
	}
	return nil
}
