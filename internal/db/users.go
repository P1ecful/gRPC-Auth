package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type UserMeta struct {
	login string
	password string
}

func ConnQuery(){
	db, err := sql.Open("sqlite", "usermeta.db")
	if err !- nil {
		log.Fatal(err)
	}

	defer db.Close()
}


func (u *UserMeta) AddMeta() {
	// db, err := sql.Open("sqlite", "usersmeta.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()
	ConnQuery()
	_, err = db.Exec("insert into Meta (login, password) values ($1, $2)", u.login, u.password)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Succesful register")
}

func (u *UserMeta) CheckMeta {
	// db, err := sql.Open("sqlite", "usersmeta.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()
	ConnQuery()
	_, err = db.QueryRow("select * form Meta where login = $1", u.login)
	if err != nil {
		log.Fatal(err)
	}
}

