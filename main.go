package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type usuario struct {
	Id    int           `json:"id"`
	Nome  string        `json:"nome"`
	Email string        `json:"email"`
	Idade sql.NullInt64 `json:"idade"`
}

func main() {
	router := mux.NewRouter()

	fmt.Println("Servidor rodando na porta -> :5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
