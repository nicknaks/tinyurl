package tinyUrl

import (
	"database/sql"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

func getPostgres() *sql.DB {
	db, err := sql.Open("pgx", "dbname=docker user=docker password=docker host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		panic("cant parse config" + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("can`t ping db" + err.Error())
	}

	db.SetMaxOpenConns(10)

	return db
}

func Init() *sqlx.DB {
	return sqlx.NewDb(getPostgres(), "psx")
}
