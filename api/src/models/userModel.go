package models

import (
	auth "api/src/Auth"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdat,omitempty"`
}

// Prepare chama os métods para validar e formatar o usuario recebido
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.formater(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("o e-mail inserido é invalido")
	}

	if step == "creation" && user.Password == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (user *User) formater(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "creation" {
		hashPassword, err := auth.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPassword)
	}

	return nil
}
