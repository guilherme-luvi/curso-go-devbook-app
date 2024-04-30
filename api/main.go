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

	// Carrgar rotas da API
	router := router.Gerar()

	// Executar aplicação na porta definida nas variaveis de ambiente
	fmt.Println("Rodando API!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
