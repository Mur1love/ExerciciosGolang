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

	// Mudando o nome e o ID do usuario pelo nome e ID - UPDATE
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Washington Estive", 1)
	stmt.Exec("Mariana Seixas", 2)

	// Deletando um usuario pelo id - DELETE
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)

}
