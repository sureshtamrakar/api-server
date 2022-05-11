package models_login

import (
	database_sqlite "github.com/sureshtamrakar/api-server/database/sqlite"
)

type Entity struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Country  string `json:"country"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

func Login(email string) (login Entity, err error) {
	row := database_sqlite.EsConn.Conn.QueryRow("SELECT email, password FROM `user` WHERE `email`=?", email)
	err = row.Scan(
		&login.Email,
		&login.Password,
	)
	return login, err
}

func Create(email, name, gender, country, password string, age int) error {
	stmt, err := database_sqlite.EsConn.Conn.Prepare("INSERT INTO user (email, name, gender, country, age, password) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(email, name, gender, country, age, password)
	if err != nil {
		return err
	}
	stmt.Close()
	return nil
}
