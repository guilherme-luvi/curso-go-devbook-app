package repositories

import (
	"api/src/models"
	"database/sql"
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
