package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// String de conexão com o MySql
	DatabaseStringConnection string

	// Porta onde a API estará rodando
	Port int

	// Secret Key é a chave que será utilizada para assinar os tokens jwt
	SecretKey []byte
)

// Função que irá inicializar as variaveis de ambiente da aplicação
func LoadEnvVars() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		// Valor padrao em caso de erro na leitura da variavel
		Port = 9000
	}

	DatabaseStringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
