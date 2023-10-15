package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var connStr = "user=p1ecful password=p1ecful dbname=Meta sslmode=disable"
var db, _ = sql.Open("postgres", connStr)

type User struct {
	UserHash string
}

func (u User) Add() string {
	db.Exec("insert into Meta (userHash) values ($1)", u.UserHash)
	defer db.Close()

	return "Succesfull"
}

func (u User) Auth() string {
	var response string
	dbMeta := User{}
	defer db.Close()

	f := db.QueryRow("select * from Meta where userHash = $1", u.UserHash)
	f.Scan(&dbMeta.UserHash)

	if dbMeta.UserHash == u.UserHash {
		response = "200 OK!"
	} else {
		response = "Error"
	}

	return response

}
