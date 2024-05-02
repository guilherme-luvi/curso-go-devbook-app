package controllers

import (
	responses "api/src/Responses"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	// Leitura do body da requisição
	requestBody, ex := ioutil.ReadAll(request.Body)
	if ex != nil {
		responses.Erro(writer, http.StatusUnprocessableEntity, ex)
		return
	}

	// Desserializar json da requisição no model de User
	var user models.User
	if ex = json.Unmarshal(requestBody, &user); ex != nil {
		responses.Erro(writer, http.StatusBadRequest, ex)
		return
	}

	// Validar usuario recebido
	if ex = user.Prepare(); ex != nil {
		responses.Erro(writer, http.StatusBadRequest, ex)
		return
	}

	// Abrir conexão com banco
	db, ex := database.Connect()
	if ex != nil {
		responses.Erro(writer, http.StatusInternalServerError, ex)
		return
	}
	defer db.Close()

	// Criar instancia do repositório de usuarios e chamar método de criação passando a variavel user
	repository := repositories.NewUsersRepository(db)
	user.ID, ex = repository.Create(user)
	if ex != nil {
		responses.Erro(writer, http.StatusInternalServerError, ex)
		return
	}

	responses.JSON(writer, http.StatusCreated, user)
}

// Busca usuarios no banco com base na query da URL
// parametro 'user' na query
// Busca por nome ou nick
func GetUsers(writer http.ResponseWriter, request *http.Request) {
	nameOrNick := strings.ToLower(request.URL.Query().Get("user"))

	db, ex := database.Connect()
	if ex != nil {
		responses.Erro(writer, http.StatusInternalServerError, ex)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, ex := repository.Get(nameOrNick)
	if ex != nil {
		responses.Erro(writer, http.StatusInternalServerError, ex)
		return
	}

	responses.JSON(writer, http.StatusOK, users)
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
