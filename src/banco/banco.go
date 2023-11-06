package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// essa funcção vai conectar com o banco de dados e retorna
func conectar() (*sql.DB, error) {

	db, erro := sql.Open("mysql", config.StringConexadoBanco)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
