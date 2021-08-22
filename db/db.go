package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoDeDados() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/LojinhaGO")

	if err != nil {
		panic(err.Error())
	}

	return db
}
