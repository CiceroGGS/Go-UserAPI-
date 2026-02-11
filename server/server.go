package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Idade int    `json:"idade"`
}

// CriarUsuario realiza a cricao de um novo usuario no banco
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoDaRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler corpo da requisição", http.StatusUnprocessableEntity)
		return
	}

	var u usuario
	if err = json.Unmarshal(corpoDaRequisicao, &u); err != nil {
		http.Error(w, "Erro ao converter usuário para struct", http.StatusBadRequest)
		return
	}
	fmt.Println(u)

	db, err := db.ConetarComBandoDeDados()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	statement, err := db.Prepare("INSERT INTO usuarios (nome, email, idade) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(u.Nome, u.Email, u.Idade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idInserido, err := insert.LastInsertId()

	w.WriteHeader(201)
	w.Write([]byte(fmt.Sprintf("Usuario criado com sucesso! id - %d |", idInserido)))
}
