package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", server.CriarUsuario).Methods(http.MethodPost)

	fmt.Println("Servidor rodando na porta -> :5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
