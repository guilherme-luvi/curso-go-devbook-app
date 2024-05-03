package controllers

import (
	responses "api/src/Responses"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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
	if ex = user.Prepare("creation"); ex != nil {
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

// Busca usuário por ID no path da requisição
func GetUserById(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)

	userID, ex := strconv.ParseUint(parameters["userId"], 10, 64) //conversao na base 10, 64 bits
	if ex != nil {
		responses.Erro(writer, http.StatusBadRequest, ex)
		return
	}

	db, ex := database.Connect()
	if ex != nil {
		responses.Erro(writer, http.StatusInternalServerError, ex)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user, ex := repository.GetById(userID)
	if ex != nil {
		responses.Erro(writer, http.StatusInternalServerError, ex)
		return
	}

	responses.JSON(writer, http.StatusOK, user)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	userID, ex := strconv.ParseUint(parameters["userId"], 10, 64)
	if ex != nil {
		responses.Erro(writer, http.StatusBadRequest, ex)
		return
	}

	requestBody, ex := ioutil.ReadAll(request.Body)
	if ex != nil {
		responses.Erro(writer, http.StatusUnprocessableEntity, ex)
		return
	}

	var user models.User
	if ex = json.Unmarshal(requestBody, &user); ex != nil {
		responses.Erro(writer, http.StatusBadRequest, ex)
		return
	}

	if ex = user.Prepare("update"); ex != nil {
		responses.Erro(writer, http.StatusBadRequest, ex)
		return
	}

	db, ex := database.Connect()
	if ex != nil {
		responses.Erro(writer, http.StatusInternalServerError, ex)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if ex = repository.Update(userID, user); ex != nil {
		responses.Erro(writer, http.StatusInternalServerError, ex)
		return
	}

	responses.JSON(writer, http.StatusOK, nil)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Deletando Usuário!"))
}
