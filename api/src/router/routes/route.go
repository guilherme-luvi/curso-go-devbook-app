package routes

import "net/http"

// o struct Route representa todas as rotas da nossa API
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}
