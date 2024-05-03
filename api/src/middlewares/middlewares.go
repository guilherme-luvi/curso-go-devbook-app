package middlewares

import (
	auth "api/src/Auth"
	responses "api/src/Responses"
	"log"
	"net/http"
)

func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Verifica se o usuario fazendo a requisição está autenticado
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Erro(w, http.StatusUnauthorized, err)
			return
		}

		nextFunction(w, r)
	}
}
