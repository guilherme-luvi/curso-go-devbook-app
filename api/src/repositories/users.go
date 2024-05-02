package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Struct que manipula os dados no banco
// Representa um respositório
type users struct {
	db *sql.DB
}

// Função que recebe uma conexão aberta com o banco
// E cria uma instacia do struct 'users' passando a conexão
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Método do struct user
// recebe um model de User e retorna um uint e um erro
// Cria um novo registro de usuario no banco
func (repository users) Create(user models.User) (uint64, error) {
	statement, ex := repository.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)
	if ex != nil {
		return 0, nil
	}
	defer statement.Close()

	result, ex := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if ex != nil {
		return 0, ex
	}

	userId, ex := result.LastInsertId()
	if ex != nil {
		return 0, nil
	}

	return uint64(userId), nil
}

// Método que busca todos os usuario que atendem ao filtro de nome ou nick
func (repository users) Get(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) //%nameOrNick%

	results, err := repository.db.Query(
		"select id, name, nick, email, createdat from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer results.Close()

	var users []models.User

	for results.Next() {
		var user models.User

		if err = results.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
