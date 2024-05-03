package controllers

import (
	auth "api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	responses "api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

// Rota para autenticação. Retorna token jwt
func Login(writer http.ResponseWriter, request *http.Request) {
	requestBody, ex := io.ReadAll(request.Body)
	if ex != nil {
		responses.Erro(writer, http.StatusUnprocessableEntity, ex)
		return
	}

	var user models.User
	if ex = json.Unmarshal(requestBody, &user); ex != nil {
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
	userRegisteredInDatabase, ex := repository.GetByEmail(user.Email)
	if ex != nil {
		responses.Erro(writer, http.StatusInternalServerError, ex)
		return
	}

	if ex = auth.VerifyPassword(userRegisteredInDatabase.Password, user.Password); ex != nil {
		responses.Erro(writer, http.StatusUnauthorized, ex)
		return
	}

	token, ex := auth.CreateToken(userRegisteredInDatabase.ID)
	if ex != nil {
		responses.Erro(writer, http.StatusInternalServerError, ex)
		return
	}

	responses.JSON(writer, http.StatusOK, token)
}
