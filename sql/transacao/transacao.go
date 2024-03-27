package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:rm@74123@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?, ?)")

	stmt.Exec(2000, "Bea")
	stmt.Exec(2001, "Lara")
	_, err = stmt.Exec(1, "Lourival") // chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
