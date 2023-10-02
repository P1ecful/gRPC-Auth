package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var db, _ = sql.Open("sqlite", "storage.db")

type User struct {
	UserHash string
}

func (u User) Add() string {
	db.Exec("insert into Meta (userHash) values ($1)", u.UserHash)
	defer db.Close()

	return "Succesfull"
}

func (u User) Auth() string {
	dbMeta := User{}
	defer db.Close()

	f := db.QueryRow("select * from Meta where userHash = $1", u.UserHash)
	f.Scan(&dbMeta.UserHash)

	return "Succesfull"

}
