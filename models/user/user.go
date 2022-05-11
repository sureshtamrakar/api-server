package models_user

import (
	database_sqlite "github.com/sureshtamrakar/api-server/database/sqlite"
)

type Entity struct {
	Id      int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Country string `json:"country"`
	Age     int    `json:"age"`
}

func GetId(email string) (value Entity, err error) {
	row := database_sqlite.EsConn.Conn.QueryRow("SELECT id FROM `user` WHERE `email`=?", email)
	err = row.Scan(
		&value.Id,
	)
	return value, err
}

func Load(id int) (value Entity, err error) {
	row := database_sqlite.EsConn.Conn.QueryRow("SELECT id,email,name,gender,country,age FROM `user` WHERE `id`=?", id)
	err = row.Scan(
		&value.Id,
		&value.Email,
		&value.Name,
		&value.Gender,
		&value.Country,
		&value.Age,
	)
	return value, err
}
