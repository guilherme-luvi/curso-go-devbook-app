package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um Router com as rotas da API configuradas
func Gerar() *mux.Router {
	router := mux.NewRouter()
	return routes.ConfigRoutes(router)
}
