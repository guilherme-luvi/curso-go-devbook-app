package auth

import (
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
	permissions["expiration"] = time.Now().Add(time.Hour * 3).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("Secret"))
}
