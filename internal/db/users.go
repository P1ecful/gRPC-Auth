package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var connStr = "user=postgres password=p1ecful dbname=storage sslmode=disable"
var db, _ = sql.Open("postgres", connStr)

type User struct {
	UserHash string
}

func (u User) Register() string {
	defer db.Close()

	if u.CheckIn() {
		db.Exec("insert into meta (userHash) values ($1)", u.UserHash)
		return "Пользователь уже зарегестрирован"
	}

	return "Успешно, вы зарегестрированы!"
}

func (u User) Auth() string {
	if !u.CheckIn() {
		return "Пользователь не найден"
	}

	return "Пользователь найден"
}

func (u User) CheckIn() bool {
	dbMeta := User{}
	defer db.Close()

	f := db.QueryRow("select * from Meta where userHash = $1", u.UserHash)
	f.Scan(&dbMeta.UserHash)

	return dbMeta.UserHash == u.UserHash
}
