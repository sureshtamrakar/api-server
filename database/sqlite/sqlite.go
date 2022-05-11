package database_sqlite

import (
	"database/sql"
	"fmt"
	"log"

	util "github.com/sureshtamrakar/api-server/util"

	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	Conn *sql.DB
}

var EsConn Connection

func init() {
	db, err := sql.Open(
		"sqlite3",
		fmt.Sprintf("./sqlite3/%s.%s", util.Yamlvalue.DBName, "db"),
	)

	if err != nil {
		log.Println("failed to create db", err)
	}
	statement, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS user  (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT  NOT NULL UNIQUE,
		name TEXT  NOT NULL,
		gender TEXT  NOT NULL,
		country TEXT  NOT NULL,
		age INTEGER NULL DEFAULT NULL,
		password TEXT  NOT NULL
		);`)
	if err != nil {
		log.Println("failed to create table", err)
	}
	statement.Exec()
	EsConn.Conn = db

}
