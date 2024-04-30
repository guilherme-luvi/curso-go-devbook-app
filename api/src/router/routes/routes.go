package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// o struct Route representa todas as rotas da nossa API
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// ConfigRoutes coloca todas as rotas dentro do router
func ConfigRoutes(router *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return router
}
