package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
)

// Recebe uma string e coloca um hash nela
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compara senha e um hash e retorna se são iguais
func VerifyPassword(hashPassword, openPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(openPassword))
}

// Cria token JWT para autenticação
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 3).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

// Extrai o ID do usuário do token
func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)

	// Método Parse obtem os claims do token recebido
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, err
	}

	// Verifica se os claims que a aplicação gera nos tokens pôde ser obtido no token informado e se o token é válido
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Retorna o ID do usuário
		userID := permissions["userId"]
		if userIDFloat, ok := userID.(float64); ok {
			return uint64(userIDFloat), nil
		}
	}

	return 0, errors.New("token inválido")
}

// Valida se o token passado na requisição é valido
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	// Método Parse obtem os claims do token recebido
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}

	// Verifica se os claims que a aplicação gera nos tokens pôde ser obtido no token informado e se o token é válido
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	// Valida Formato Bearer xyzw....
	if len(strings.Split(token, " ")) == 2 {
		// Retorna apenas o token sem a palavra Bearer
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	// Verifica se o token informado utilizou o memso método de assinatura que a aplicação utiliza para criar tokens
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	// Retorna a secret key após a verificação
	return config.SecretKey, nil
}
