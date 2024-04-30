package router

import "github.com/gorilla/mux"

// Gerar vai retornar um Router com as rotas da API configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
