package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

// Abre conexão com o banco de dados
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.DatabaseStringConnection)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		// Fechar conexão em caso de erro
		db.Close()
		return nil, erro
	}

	return db, nil
}
