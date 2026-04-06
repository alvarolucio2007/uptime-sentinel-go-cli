// Package database conecta à base de dados
package database

import (
	"errors"
	"fmt"
	"log"
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
		models.ErroAberturaPostgres.Log(err)
	}
	if err = DB.AutoMigrate(&models.ModeloLink{}); err != nil {
		models.ErroConexaoPostgres.Log(err)
	}
	fmt.Println("Postgres conectado")
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
	resultado := DB.Delete(&models.ModeloLink{}, ID)
	if resultado.Error != nil {
		models.ErroDeletePostgres.Log(resultado.Error)
		return resultado.Error
	}

	if resultado.RowsAffected == 0 {
		fmt.Println("ID inexistente, nada deletado.")
	}
	return nil
}

func BuscarEntradaPostgres(ID string) (*models.ModeloLink, error) {
	var linkEncontrado models.ModeloLink

	res := DB.Where("ID LIKE ?", "%"+ID+"%").First(&linkEncontrado)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			models.ErroBuscaPostgres.Log(res.Error)
			return nil, res.Error
		}
		log.Print(models.ErroBuscaPostgresNEncontrado.Mensagem)
		return nil, res.Error
	}
	return &linkEncontrado, nil
}
