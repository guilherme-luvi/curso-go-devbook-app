package auth

import "golang.org/x/crypto/bcrypt"

// Recebe uma string e coloca um hash nela
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compara senha e um hash e retorna se são iguais
func VerifyPassword(hashPassword, openPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(openPassword))
}
