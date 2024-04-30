package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// Função main que executa a aplicação
func main() {
	// Carregar variáveis de ambiente
	config.LoadEnvVars()

	// Carregar rotas da API
	router := router.Gerar()

	// Executar aplicação na porta definida nas variaveis de ambiente
	fmt.Printf("API executando na porta %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
