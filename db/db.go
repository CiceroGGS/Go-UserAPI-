package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// ConetarComBandoDeDados => Funcao para conectar com o banco mysql com a connection string
func ConetarComBandoDeDados() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	connectionString := os.Getenv("DB_URL")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Erro ao abrir conexao com banco de dados")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Erro ao conectar com o banco de dados")
		return nil, err
	}

	return db, nil
}
