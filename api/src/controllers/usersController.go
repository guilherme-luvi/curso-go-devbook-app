package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	// Leitura do body da requisição
	requestBody, ex := ioutil.ReadAll(request.Body)
	if ex != nil {
		log.Fatal(ex)
	}

	// Desserializar json da requisição no model de User
	var user models.User
	if ex = json.Unmarshal(requestBody, &user); ex != nil {
		log.Fatal(ex)
	}

	// Abrir conexão com banco
	db, ex := database.Connect()
	if ex != nil {
		log.Fatal(ex)
	}

	// Criar instancia do repositório de usuarios e chamar método de criação passando a variavel user
	repository := repositories.NewUsersRepository(db)
	userId, ex := repository.Create(user)
	if ex != nil {
		log.Fatal(ex)
	}

	writer.Write([]byte(fmt.Sprintf("User created: %d", userId)))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuarios!"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuario por id!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Editando Usuario!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
